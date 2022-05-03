/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package new

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/monkeswag33/noter-go/db"
	"github.com/monkeswag33/noter-go/errordef"
	"github.com/monkeswag33/noter-go/global"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// noteCmd represents the note command
var noteCmd = &cobra.Command{
	Use:   "note <note name>",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		database = db.Database
		username, _ := cmd.Flags().GetString("user")
		body, _ := cmd.Flags().GetString("body")
		note, err := createNote(args, username, body)
		if err != nil {
			logrus.Fatal(err)
		}
		if err := insertNote(note); err != nil {
			logrus.Fatal(err)
		}
		fmt.Printf("Created note %q\n", note.Name)
	},
}

func createNote(args []string, username string, body string) (*db.Note, error) {
	var name string
	if len(args) == 1 {
		logrus.Debug("Note name given as argument, using it")
		name = args[0]
		if err := newNoteValidateNoteName(name); err != nil {
			return nil, err
		}
	} else {
		name = global.Prompt(promptui.Prompt{}, "Note name:", newNoteValidateNoteName)
	}
	logrus.Debug("Note name passed validation")
	if len(username) == 0 {
		logrus.Debug("Username not given as parameter, prompting for it")
		username = global.Prompt(promptui.Prompt{}, "User note belongs to?", newNoteValidateUsername)
	} else if err := newNoteValidateUsername(username); err != nil {
		return nil, err
	}
	logrus.Debug("Username passed validation")
	if len(body) == 0 {
		logrus.Debug("Body not given as parameter, prompting for it")
		body = getBody()
		logrus.Debug("Got body of note")
	} else {
		logrus.Debug("Body given as parameter, using it")
	}
	var note db.Note = db.Note{
		Name: name,
		Body: body,
		User: db.User{
			Username: username,
		},
	}
	logrus.Info("Created note")
	return &note, nil
}

func insertNote(note *db.Note) error {
	if err := database.CreateNote(note); err != nil {
		return err
	}
	logrus.Info("Inserted note")
	return nil
}

func init() {
	NewCmd.AddCommand(noteCmd)
	noteCmd.Flags().StringP("user", "u", "", "User that the note will be added to")
	noteCmd.Flags().StringP("body", "b", "", "Body of note")
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

func newNoteValidateNoteName(noteName string) error {
	if len(noteName) < 5 {
		return errordef.ErrNoteNameTooShort
	}
	exists, err := database.CheckNoteExists(db.Note{
		Name: noteName,
	})
	if err != nil {
		logrus.Fatal(err)
	}
	if exists {
		return errordef.ErrNoteAlreadyExists
	}
	return nil
}

func newNoteValidateUsername(username string) error {
	exists, err := database.CheckUserExists(db.User{
		Username: username,
	})
	if err != nil {
		logrus.Fatal(err)
	}
	if !exists {
		return errordef.ErrUserDoesntExist
	}
	return nil
}
