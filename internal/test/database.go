package test

import (
	"fmt"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewTestDB(t *testing.T, models ...any) *gorm.DB {
	t.Helper()
	dsn := fmt.Sprintf("file:%s?mode=memory&cache=shared", t.Name())

	testDB, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = testDB.AutoMigrate(
		models...,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		sqlDB, err := testDB.DB()
		if err != nil {
			t.Fatal(err)
		}

		err = sqlDB.Close()
		if err != nil {
			t.Fatal(err)
		}
	})

	return testDB
}
