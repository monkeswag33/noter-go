package describe

import (
	"errors"

	"github.com/manifoldco/promptui"
	"github.com/monkeswag33/noter-go/db"
	"github.com/monkeswag33/noter-go/global"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var noteCmd = &cobra.Command{
	Use:   "note <note name>",
	Short: "Describe a specific note",
	Long:  "Get all the information about a specific note",
	Run: func(cmd *cobra.Command, args []string) {
		var noteName string
		if len(args) != 1 {
			logrus.Debug("Note name not found, prompting")
			noteName = global.Prompt(promptui.Prompt{}, "Note name:", describeNoteValidateNote)
		} else {
			noteName = args[0]
			if err := describeNoteValidateNote(noteName); err != nil {
				logrus.Fatal(err)
			}
		}
		logrus.Debug("Note name passed validation")
		var note db.Note = db.GetNotes("", 0, noteName)[0]
		logrus.Debug("Received note")
		printFormatted(note.Name, []string{"Id", "Name", "Body", "Owner"}, []interface{}{
			note.ID,
			note.Name,
			note.Body,
			note.User.Username,
		})
	},
}

func describeNoteValidateNote(noteName string) error {
	if !db.CheckNoteExists(noteName) {
		return errors.New("note does not exist")
	}
	return nil
}

func init() {
	DescribeCmd.AddCommand(noteCmd)
}
