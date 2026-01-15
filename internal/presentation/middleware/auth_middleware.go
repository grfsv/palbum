package middleware

import (
	"palbum/internal/application/service"
	"palbum/internal/domain/commons"
	"palbum/internal/domain/user"
	"palbum/internal/presentation/utils"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	tokenService service.TokenService
	errorHandler *utils.ErrorHandler
}

func NewAuthMiddleware(tokenService service.TokenService, errorHandler *utils.ErrorHandler) *AuthMiddleware {
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

func MustGetUserUUID(ctx *gin.Context) user.UUID {
	uuid, exists := ctx.Get("userUUID")
	if !exists {
		panic("userUUID not found in context")
	}

	uuidStr, ok := uuid.(string)
	if !ok {
		panic("invalid userUUID in context")
	}

	userUUID, err := user.NewUUIDFromString(uuidStr)
	if err != nil {
		panic("invalid userUUID in context")
	}

	return userUUID
}

func trimPrefix(tokenString string) (string, bool) {
	prefix := "Bearer "

	if len(tokenString) < len(prefix) || tokenString[:len(prefix)] != prefix {
		return "", false
	}

	return tokenString[len(prefix):], true
}
