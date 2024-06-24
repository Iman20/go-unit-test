package config

import (
	"go-unit-test/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	var err error

	dbUrl := "postgres://localhost:5432/otp"

	dbCon, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	dbCon.AutoMigrate(models.User{})

	if err != nil {
		panic(err)
	}

	return dbCon
}

func InitMockDB() *gorm.DB {
	var err error

	dbUrl := "file::memory:?cache=shared"

	dbCon, err := gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})

	dbCon.AutoMigrate(models.User{})

	if err != nil {
		panic(err)
	}

	return dbCon
}
