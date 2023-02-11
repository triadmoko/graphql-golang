package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/triadmoko/grahpql-golang/graph"
	"github.com/triadmoko/grahpql-golang/graph/resolver"
)

func main() {
	// Setting up Gin
	r := gin.Default()
	r.POST("/query", func(c *gin.Context) {
		h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}}))
		h.ServeHTTP(c.Writer, c.Request)
	})

	r.GET("/", func(c *gin.Context) {
		h := playground.Handler("GraphQL", "/query")
		h.ServeHTTP(c.Writer, c.Request)
	})

	r.Run(":8081")
}
