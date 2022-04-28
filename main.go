/*
Copyright © 2022 NAME HERE ishan.karmakar24@gmail.com

*/
package main

import (
	"github.com/monkeswag33/noter-go/cmd"
	"github.com/monkeswag33/noter-go/db"
	"github.com/monkeswag33/noter-go/global"
)

func main() {
	global.SetupViper()
	db.InitDB(global.SetLogLevel()) // SetLogLevel returns gorm log level, which is passed to initdb
	cmd.Execute()
}
