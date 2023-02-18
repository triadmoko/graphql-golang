package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"time"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/triadmoko/grahpql-golang/config"
	"github.com/triadmoko/grahpql-golang/handler"
	"github.com/triadmoko/grahpql-golang/helpers"
	"github.com/triadmoko/grahpql-golang/injector"
	"github.com/triadmoko/grahpql-golang/middleware"
)

func main() {
	// load config
	conf, err := config.InitConfig()
	if err != nil {
		log.Fatal("Error Load Config", err.Error())
		return
	}
	// Setting up Gin
	r := gin.Default()
	r.POST("/query",
		middleware.Middleware(),
		ginBodyLogMiddleware(conf),
		handler.GraphqlHandler(
			injector.NewInitInjector(conf),
		))

	r.GET("/", func(c *gin.Context) {
		h := playground.Handler("GraphQL", "/query")
		h.ServeHTTP(c.Writer, c.Request)
	})

	r.Run(":8081")
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func ginBodyLogMiddleware(conf *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		sess, _ := c.Request.Context().Value("sess").(*helpers.MetaToken)

		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		body, _ := ioutil.ReadAll(tee)
		c.Request.Body = ioutil.NopCloser(&buf)

		start := time.Now().UTC().UnixMilli()
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}

		c.Writer = blw
		c.Next()

		finish := time.Now().UTC().UnixMilli()
		log := Logs{
			Path:      c.Request.RequestURI,
			Duration:  finish - start,
			Response:  blw.body.String(),
			Timestamp: time.Now(),
			User:      sess,
			Request:   string(body),
		}
		coll := conf.Mongo.Collection("log")
		result, err := coll.InsertOne(context.TODO(), log)
		if err != nil {
			conf.Logger.Error("Error insert log")
		}
		fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	}
}

type Logs struct {
	Path      string
	User      *helpers.MetaToken
	Duration  int64
	Request   string
	Response  string
	Timestamp time.Time
}
