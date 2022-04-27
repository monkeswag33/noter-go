package get

import (
	"fmt"
	"os"

	"github.com/monkeswag33/noter-go/db"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Command to get all users",
	Long:  "Command to get all users",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		id, _ := cmd.Flags().GetInt("id")
		var table *tablewriter.Table = tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Id", "Name"})
		var users []db.User = db.GetUsers(username, id)
		for _, user := range users {
			var stringrepr []string = []string{fmt.Sprint(user.ID), user.Username}
			table.Append(stringrepr)
		}
		table.Render()
	},
}

func init() {
	GetCmd.AddCommand(usersCmd)

	usersCmd.Flags().StringP("username", "u", "", "Search for specific username")
	usersCmd.Flags().IntP("id", "i", 0, "Search for specific id")
}
