package middleware

import (
	"github.com/gin-gonic/gin"
)

func TraceID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// requestid := requestid.Get(c)
		// ctx := key.WithRequestID(c.Request.Context(), requestid)
		// c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
