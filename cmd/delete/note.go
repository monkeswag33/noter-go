package delete

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/monkeswag33/noter-go/db"
	"github.com/monkeswag33/noter-go/errordef"
	"github.com/monkeswag33/noter-go/prompt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var noteCmd = &cobra.Command{
	Use:   "note <note name>",
	Short: "Delete note",
	Long:  "Delete a note by name",
	Run: func(cmd *cobra.Command, args []string) {
		database = db.Database
		var noteName string
		if len(args) == 1 {
			logrus.Debug("Note name recognized as argument, using it")
			noteName = args[0]
			if err := deleteNoteValidateNoteName(noteName); err != nil {
				logrus.Fatal(err)
			}
		} else {
			logrus.Debug("Note name not found, prompting")
			noteName = prompt.Prompt(promptui.Prompt{}, "Note name:", deleteNoteValidateNoteName)
		}
		logrus.Debug("Username passed validation")
		if err := database.DeleteNote(db.Note{
			Name: noteName,
		}); err != nil {
			logrus.Fatal(err)
		}
		logrus.Info("Deleted note")
		fmt.Printf("Deleted note %q\n", noteName)
	},
}

func init() {
	DeleteCmd.AddCommand(noteCmd)
}

func deleteNoteValidateNoteName(noteName string) error {
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
