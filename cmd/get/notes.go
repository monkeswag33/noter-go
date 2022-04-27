package get

import (
	"fmt"
	"os"

	"github.com/monkeswag33/noter-go/db"
	"github.com/olekukonko/tablewriter"
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
		var table *tablewriter.Table = tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Id", "Name", "Owner"})
		for _, user := range notes {
			var stringrepr []string = []string{fmt.Sprint(user.ID), user.Name, user.User.Username}
			table.Append(stringrepr)
		}
		table.Render()
	},
}

func init() {
	GetCmd.AddCommand(notesCmd)

	notesCmd.Flags().StringP("owner", "o", "", "Filter by note owner")
	notesCmd.Flags().IntP("id", "i", 0, "Filter by note id")
	notesCmd.Flags().StringP("name", "n", "", "Filter by note name")
}
