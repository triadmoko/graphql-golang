package handler

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/triadmoko/grahpql-golang/graph"
	"github.com/triadmoko/grahpql-golang/graph/resolver"
	"github.com/triadmoko/grahpql-golang/injector"
)

func GraphqlHandler(injection *injector.Injector) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	conf := graph.Config{
		Resolvers: &resolver.Resolver{
			User: injection.User,
		},
	}
	exec := graph.NewExecutableSchema(conf)
	h := handler.GraphQL(exec)
	return func(c *gin.Context) { h.ServeHTTP(c.Writer, c.Request) }
}
