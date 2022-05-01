package global

import (
	database "github.com/monkeswag33/noter-go/db"
)

var DB *database.DB

func SetupDB() {
	DB = &database.DB{
		LogLevel: SetLogLevel(),
	}
	DB.Init()
}

func ShutdownDB() {
	DB.Close()
	DB = nil
}
