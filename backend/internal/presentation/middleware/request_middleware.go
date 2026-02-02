package middleware

import (
	"bytes"
	"io"
	"palbum/internal/utils/log"

	"github.com/gin-gonic/gin"
)

type RequestMiddleware struct {
	logger *log.Log
}

func NewRequestMiddleware(logger *log.Log) *RequestMiddleware {
	return &RequestMiddleware{
		logger: logger,
	}
}

func (m *RequestMiddleware) LogRequests() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery

		var bodyBytes []byte
		if ctx.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(ctx.Request.Body)

			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		m.logger.Info("Incoming request",
			"method", method,
			"path", path,
			"query", query,
			"headers", ctx.Request.Header,
			"body", string(bodyBytes),
		)

		ctx.Next()
	}
}
