package global

import (
	database "github.com/monkeswag33/noter-go/db"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitTesterDB() *database.DB {
	var location string = "file::memory:?cache=shared" // Specify location to be in-memory
	gormDB, err := gorm.Open(sqlite.Open(location), &gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
	}
	var db database.DB = database.DB{
		LogLevel: "warn",
		DB:       gormDB,
	}
	db.Init()
	return &db
}
