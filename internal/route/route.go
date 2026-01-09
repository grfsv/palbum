package route

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"remind_map/internal/presentation"
	"remind_map/internal/presentation/middleware"
	"strings"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/dig"
)

type handlerParams struct {
	dig.In

	User *presentation.UserHandler
}

type ServerConfig struct {
	Port string
	Mode string
}

func NewHandler(params handlerParams, cfg *ServerConfig, authMiddleware *middleware.AuthMiddleware) *gin.Engine {
	gin.SetMode(cfg.Mode)

	engine := gin.Default()
	engine.Use(requestid.New())
	engine.Use(requestLogger())

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

			if name == "-" {
				return ""
			}

			return name
		})
	}

	routing(engine, params, authMiddleware)

	return engine
}

func routing(engine *gin.Engine, params handlerParams, middleware *middleware.AuthMiddleware) {
	v1engine := engine.Group("/v1")
	auth := v1engine.Group("/auth")
	{
		auth.POST("/signup", params.User.SignUp)
		auth.POST("/login", params.User.Login)
		auth.POST("/logout", params.User.Logout)
		auth.POST("/refresh", params.User.Refresh)
	}

	private := v1engine.Group("/")
	private.Use(middleware.RequireAuth())
	{

	}
}

func requestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)

			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		fmt.Println("---------------- REQUEST LOG START ----------------")
		fmt.Printf("Method: %s\n", method)
		fmt.Printf("Path:   %s\n", path)
		fmt.Printf("Query:  %s\n", query)

		fmt.Println("Headers:")
		for k, v := range c.Request.Header {
			fmt.Printf("  %s: %s\n", k, v)
		}

		fmt.Println("Body:")
		if len(bodyBytes) > 0 {
			fmt.Println(string(bodyBytes))
		} else {
			fmt.Println("(empty)")
		}
		fmt.Println("---------------- REQUEST LOG END ------------------")

		c.Next()
	}
}
