package main

import (
	"log"
	"net/http"
	"palbum/internal/dependencies"
	"palbum/internal/route"

	"github.com/gin-gonic/gin"
)

func main() {
	// DIコンテナの構築
	container, err := dependencies.InitContainer()
	if err != nil {
		panic(err)
	}

	// サーバーの起動
	// Ginのモード設定
	err = container.Invoke(
		func(engin *gin.Engine, cfg *route.ServerConfig) error {
			// ヘルスチェックエンドポイント
			engin.GET("/health", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"status":  "ok",
					"message": "palbum Map API is running started",
				})
			})

			// サーバー起動
			err := engin.Run(":" + cfg.Port)
			if err != nil {
				log.Fatalf("Failed to start server: %v", err)
			}

			return nil
		})
	if err != nil {
		log.Fatalf("Failed to invoke container: %v", err)
	}
}
