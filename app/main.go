package main

import (
	"log"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/triadmoko/grahpql-golang/config"
	"github.com/triadmoko/grahpql-golang/handler"
	"github.com/triadmoko/grahpql-golang/injector"
)

func main() {
	// load config
	conf, err := config.InitConfig()
	if err != nil {
		log.Fatal("Error Load Config", err.Error())
		return
	}
	user := injector.NewInjectorUser(conf)
	// Setting up Gin
	r := gin.Default()
	r.POST("/query", handler.GraphqlHandler(user))

	r.GET("/", func(c *gin.Context) {
		h := playground.Handler("GraphQL", "/query")
		h.ServeHTTP(c.Writer, c.Request)
	})

	r.Run(":8081")
}
