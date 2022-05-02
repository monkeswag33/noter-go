package global

import (
	database "github.com/monkeswag33/noter-go/db"
	"github.com/monkeswag33/noter-go/types"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *database.DB

func SetupDB(logLevel types.LogLevelParams) {
	DB = &database.DB{
		LogLevel: logLevel,
	}
	DB.Init()
}

func ShutdownDB() {
	DB.Close()
	DB = nil
}

func InitTesterDB() *database.DB {
	var location string = "file::memory:?cache=shared" // Specify location to be in-memory
	gormDB, err := gorm.Open(sqlite.Open(location), &gorm.Config{})
	if err != nil {
		logrus.Fatal(err)
	}
	var db database.DB = database.DB{
		LogLevel: types.LogLevelParams{
			LogLevel:     "warn",
			GormLogLevel: "warn",
		},
		DB: gormDB,
	}
	db.Init()
	return &db
}
