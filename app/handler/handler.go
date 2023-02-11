package handler

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/triadmoko/grahpql-golang/graph"
	"github.com/triadmoko/grahpql-golang/graph/resolver"
	"github.com/triadmoko/grahpql-golang/service/user"
)

func GraphqlHandler(us user.UserServices) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	conf := graph.Config{
		Resolvers: &resolver.Resolver{
			User: us,
		},
	}
	exec := graph.NewExecutableSchema(conf)
	h := handler.GraphQL(exec)
	return func(c *gin.Context) { h.ServeHTTP(c.Writer, c.Request) }
}
