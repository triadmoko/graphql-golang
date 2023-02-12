package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/triadmoko/grahpql-golang/helpers"
)

// SetUserContext set context if user has valid token otherwise none
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestToken := c.Request.Header.Get("Authorization")
		if len(requestToken) == 0 || requestToken == "" {
			c.Next()
			return
		}
		splitToken := strings.Split(requestToken, " ")
		result := splitToken[1]
		token, err := helpers.VerifyTokenHeader("JWT_SECRET", result)
		if err != nil {
			c.Next()
			return
		}
		c.Set("sess", &token.Claims)
		c.Request = setToContext(c, "sess", &token.Claims)
		c.Next()
	}
}

// Set context key-value pair
func setToContext(c *gin.Context, key interface{}, value interface{}) *http.Request {
	return c.Request.WithContext(context.WithValue(c.Request.Context(), key, value))
}
