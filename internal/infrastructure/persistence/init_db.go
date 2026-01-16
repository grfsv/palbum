package persistence

import (
	"fmt"

	"github.com/cockroachdb/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

func InitDB(cfg *DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return database, nil
}
