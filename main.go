/*
Copyright © 2022 NAME HERE ishan.karmakar24@gmail.com

*/
package main

import (
	"github.com/monkeswag33/noter-go/cmd"
	"github.com/monkeswag33/noter-go/global"
)

func main() {
	global.SetupViper()
	global.SetupDB()
	cmd.Execute()
	global.DB.Close()
}
