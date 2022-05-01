package describe

import (
	"github.com/manifoldco/promptui"
	"github.com/monkeswag33/noter-go/db"
	"github.com/monkeswag33/noter-go/errordef"
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
		notes, err := database.GetNotes(db.Note{
			Name: noteName,
		})
		if err != nil {
			logrus.Fatal(err)
		}
		var note db.Note = notes[0]
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
	exists, err := database.CheckNoteExists(db.Note{
		Name: noteName,
	})
	if err != nil {
		logrus.Fatal(err)
	}
	if !exists {
		return errordef.ErrNoteDoesntExist
	}
	return nil
}

func init() {
	DescribeCmd.AddCommand(noteCmd)
}
