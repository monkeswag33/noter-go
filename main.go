/*
Copyright © 2022 NAME HERE ishan.karmakar24@gmail.com

*/
package main

import (
	"github.com/monkeswag33/noter-go/cmd"
	"github.com/monkeswag33/noter-go/db"
	"github.com/spf13/viper"
)

func setupViper() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.AutomaticEnv()
	viper.SetDefault("LOG_LEVEL", "warn")
}

func main() {
	setupViper()
	db.InitDB(SetLogLevel()) // SetLogLevel returns gorm log level, which is passed to initdb
	cmd.Execute()
}
