package route

import (
	"remind_map/internal/presentation"
	"remind_map/internal/presentation/middleware"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
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
	engine.Use(middleware.TraceID())

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
