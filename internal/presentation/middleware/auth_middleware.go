package middleware

import (
	"palbum/internal/application/service"
	"palbum/internal/domain/commons"
	"palbum/internal/presentation"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	tokenService service.TokenService
	errorHandler *presentation.ErrorHandler
}

func NewAuthMiddleware(tokenService service.TokenService, errorHandler *presentation.ErrorHandler) *AuthMiddleware {
	return &AuthMiddleware{
		tokenService: tokenService,
		errorHandler: errorHandler,
	}
}

func (m *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		tokenString, ok := trimPrefix(tokenString)
		if !ok {
			m.errorHandler.HandleError(ctx, commons.NewUnAuthorizedError())
			ctx.Abort()

			return
		}

		userUUID, err := m.tokenService.ConfirmAccessToken(tokenString)
		if err != nil {
			m.errorHandler.HandleError(ctx, err)
			ctx.Abort()

			return
		}

		ctx.Set("userUUID", userUUID)
		ctx.Next()
	}
}

func trimPrefix(tokenString string) (string, bool) {
	prefix := "Bearer "

	if len(tokenString) < len(prefix) || tokenString[:len(prefix)] != prefix {
		return "", false
	}

	return tokenString[len(prefix):], true
}
