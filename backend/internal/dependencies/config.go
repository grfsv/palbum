package dependencies

import (
	"os"
	"palbum/internal/infrastructure/persistence"
	"palbum/internal/infrastructure/security"
	"palbum/internal/route"

	"github.com/cockroachdb/errors"
	"go.uber.org/dig"
)

type Config struct {
	dig.Out

	DBConfig     *persistence.DBConfig
	JWTConfig    *security.JWTConfig
	ServerConfig *route.ServerConfig
}

func NewConfig() (Config, error) {
	var err error

	// エラーチェック付きの取得関数を定義
	get := func(key string) string {
		if err != nil {
			return "" // 既にエラーがあれば何もしない
		}

		value := os.Getenv(key)
		if value == "" {
			err = errors.WithStack(errors.Newf("environment variable %s is empty", key))

			return ""
		}

		return value
	}

	cfg := Config{
		Out: dig.Out{},
		DBConfig: &persistence.DBConfig{
			User: get("DB_USER"),
			Pass: get("DB_PASSWORD"),
			Host: get("DB_HOST"),
			Port: get("DB_PORT"),
			Name: get("DB_NAME"),
		},
		JWTConfig: &security.JWTConfig{
			AccessTokenSecret:  get("ACCESS_TOKEN_SECRET"),
			RefreshTokenSecret: get("REFRESH_TOKEN_SECRET"),
		},
		ServerConfig: &route.ServerConfig{
			Port: get("PORT"),
			Mode: get("GIN_MODE"),
		},
	}

	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
