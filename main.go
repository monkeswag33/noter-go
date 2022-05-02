/*
Copyright © 2022 NAME HERE ishan.karmakar24@gmail.com

*/
package main

import (
	"github.com/monkeswag33/noter-go/cmd"
	"github.com/monkeswag33/noter-go/global"
	"github.com/monkeswag33/noter-go/types"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := global.SetupViper(); err != nil {
		logrus.Fatal(err)
	}
	var logLevel types.LogLevelParams = global.ParseLogLevel()
	if err := global.SetLogLevel(logLevel); err != nil {
		logrus.Fatal(err)
	}
	global.SetupDB(logLevel)
	cmd.Execute()
	global.ShutdownDB()
}
