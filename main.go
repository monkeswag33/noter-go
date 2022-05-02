/*
Copyright © 2022 NAME HERE ishan.karmakar24@gmail.com

*/
package main

import (
	"github.com/monkeswag33/noter-go/cmd"
	"github.com/monkeswag33/noter-go/global"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := global.SetupViper(); err != nil {
		logrus.Fatal(err)
	}
	global.SetupDB()
	cmd.Execute()
	global.ShutdownDB()
}
