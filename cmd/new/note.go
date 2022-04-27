/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package new

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/monkeswag33/noter-go/db"
	"github.com/monkeswag33/noter-go/global"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// noteCmd represents the note command
var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		username, _ := cmd.Flags().GetString("user")
		if len(name) == 0 {
			logrus.Debug("Note name not given as parameter, prompting for it")
			name = global.Prompt(promptui.Prompt{}, "Note name:", newNoteValidateNoteName)
		} else if err := newNoteValidateNoteName(name); err != nil {
			logrus.Fatal(err)
		}
		logrus.Debug("Note name passed validation")
		if len(username) == 0 {
			logrus.Debug("Username not given as parameter, prompting for it")
			username = global.Prompt(promptui.Prompt{}, "User note belongs to?", newNoteValidateUsername)
		} else if err := newNoteValidateUsername(username); err != nil {
			logrus.Fatal(err)
		}
		logrus.Debug("Username passed validation")
		var body string = getBody()
		logrus.Debug("Got body of note")
		note, err := db.CreateNote(name, body, username)
		logrus.Info("Created note")
		if err != nil {
			logrus.Fatal(err)
		}
		fmt.Printf("Created note %q\n", note.Name)
	},
}

func newNoteValidateNoteName(noteName string) error {
	if len(noteName) < 5 {
		return errors.New("note name must be at least 5 characters long")
	}
	if db.CheckNoteExists(noteName) {
		return errors.New("note already exists")
	}
	return nil
}

func newNoteValidateUsername(username string) error {
	if !db.CheckUserExists(username) {
		return errors.New("user does not exist")
	}
	return nil
}

func init() {
	NewCmd.AddCommand(noteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// noteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	noteCmd.Flags().StringP("name", "n", "", "Name of the note")
	noteCmd.Flags().StringP("user", "u", "", "User that the note will be added to")
}

func getBody() string {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var lines string
	fmt.Println("Enter body (CTRL+] when finished): ")
	for {
		scanner.Scan()
		var text string = scanner.Text()
		if len(text) == 1 && text[0] == '\x1D' {
			break
		}
		lines += text + "\n"
	}
	lines = strings.TrimSuffix(lines, "\n")
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
	return lines
}
