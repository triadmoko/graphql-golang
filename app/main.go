package main

import (
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/triadmoko/grahpql-golang/config"
	"github.com/triadmoko/grahpql-golang/graph"
	"github.com/triadmoko/grahpql-golang/graph/resolver"
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
	r.POST("/query", func(c *gin.Context) {
		h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
			Resolvers: &resolver.Resolver{
				Config: conf,
			},
		}))
		h.ServeHTTP(c.Writer, c.Request)
	})

	r.GET("/", func(c *gin.Context) {
		h := playground.Handler("GraphQL", "/query")
		h.ServeHTTP(c.Writer, c.Request)
	})

	r.Run(":8081")
}
