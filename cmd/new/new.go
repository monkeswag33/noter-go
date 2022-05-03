/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package new

import (
	"github.com/monkeswag33/noter-go/db"
	"github.com/spf13/cobra"
)

var database *db.DB

// newCmd represents the new command
var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create object",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}
