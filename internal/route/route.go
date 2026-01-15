package route

import (
	"palbum/internal/presentation"
	"palbum/internal/presentation/middleware"
	"reflect"
	"strings"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/dig"
)

type handlerParams struct {
	dig.In

	User   *presentation.UserHandler
	Friend *presentation.FriendHandler
}

type middlewareParams struct {
	dig.In

	AuthMiddleware    *middleware.AuthMiddleware
	RequestMiddleware *middleware.RequestMiddleware
}

type ServerConfig struct {
	Port string
	Mode string
}

func NewHandler(params handlerParams, cfg *ServerConfig, middlewareParams middlewareParams) *gin.Engine {
	gin.SetMode(cfg.Mode)

	engine := gin.Default()
	engine.Use(requestid.New())
	engine.Use(middlewareParams.RequestMiddleware.LogRequests())

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			tag := fld.Tag.Get("json")
			name, _, _ := strings.Cut(tag, ",")

			if name == "-" {
				return ""
			}

			return name
		})
	}

	routing(engine, params, middlewareParams.AuthMiddleware)

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
		friend := private.Group("/friend")
		{
			friend.GET("/code", params.Friend.GetFriendCode)
			friend.POST("/request/:friend_code", params.Friend.RequestFriend)
			friend.PATCH("/request/:request_uuid", params.Friend.UpdateRequestStatus)
		}
	}
}
