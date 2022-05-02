package global

import (
	database "github.com/monkeswag33/noter-go/db"
	"github.com/monkeswag33/noter-go/types"
)

var DB *database.DB

func SetupDB() {
	var logLevel types.LogLevelParams = ParseLogLevel()
	DB = &database.DB{
		LogLevel: logLevel,
	}
	DB.Init()
}

func ShutdownDB() {
	DB.Close()
	DB = nil
}
