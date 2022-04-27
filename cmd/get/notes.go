package get

import (
	"fmt"
	"os"

	"github.com/monkeswag33/noter-go/db"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "Command to get all notes",
	Long:  "Command to get all notes",
	Run: func(cmd *cobra.Command, args []string) {
		owner, _ := cmd.Flags().GetString("owner")
		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")
		var notes []db.Note = db.GetNotes(owner, id, name)
		logrus.Debug("Retrieved list of notes")
		var table *tablewriter.Table = tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Id", "Name", "Owner"})
		logrus.Trace("Set headers of table")
		for _, user := range notes {
			var stringrepr []string = []string{fmt.Sprint(user.ID), user.Name, user.User.Username}
			table.Append(stringrepr)
			logrus.Trace("Added row to table")
		}
		logrus.Debug("Rendering table...")
		table.Render()
	},
}

func init() {
	GetCmd.AddCommand(notesCmd)

	notesCmd.Flags().StringP("owner", "o", "", "Filter by note owner")
	notesCmd.Flags().IntP("id", "i", 0, "Filter by note id")
	notesCmd.Flags().StringP("name", "n", "", "Filter by note name")
}
