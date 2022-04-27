/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	deleteCmd "github.com/monkeswag33/noter-go/cmd/delete"
	describeCmd "github.com/monkeswag33/noter-go/cmd/describe"
	getCmd "github.com/monkeswag33/noter-go/cmd/get"
	newCmd "github.com/monkeswag33/noter-go/cmd/new"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "noter",
	Short: "A CLI to manage notes",
	Long: `This is a CLI to manage notes and users. Each user has their own notes, so you can organize
your notes`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(newCmd.NewCmd)
	rootCmd.AddCommand(getCmd.GetCmd)
	rootCmd.AddCommand(describeCmd.DescribeCmd)
	rootCmd.AddCommand(deleteCmd.DeleteCmd)
}
