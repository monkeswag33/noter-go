package delete

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/monkeswag33/noter-go/db"
	"github.com/monkeswag33/noter-go/global"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var noteCmd = &cobra.Command{
	Use:   "note <note name>",
	Short: "Delete note",
	Long:  "Delete a note by name",
	Run: func(cmd *cobra.Command, args []string) {
		var noteName string
		if len(args) == 1 {
			logrus.Debug("Note name recognized as argument, using it")
			noteName = args[0]
			if err := deleteNoteValidateNoteName(noteName); err != nil {
				logrus.Fatal(err)
			}
		} else {
			logrus.Debug("Note name not found, prompting")
			noteName = global.Prompt(promptui.Prompt{}, "Note name:", deleteNoteValidateNoteName)
		}
		logrus.Debug("Username passed validation")
		var deletedNote db.Note = db.DeleteNote(noteName)
		logrus.Info("Deleted note")
		fmt.Printf("Deleted note %q\n", deletedNote.Name)
	},
}

func init() {
	DeleteCmd.AddCommand(noteCmd)
}

func deleteNoteValidateNoteName(noteName string) error {
	if !db.CheckNoteExists(noteName) {
		return errors.New("note does not exist")
	}
	return nil
}
