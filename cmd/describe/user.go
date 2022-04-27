package describe

import (
	"errors"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/monkeswag33/noter-go/db"
	"github.com/monkeswag33/noter-go/global"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user <username>",
	Short: "Describe a specific user",
	Long:  "Get all the information about a specific user",
	Run: func(cmd *cobra.Command, args []string) {
		var username string
		if len(args) == 1 {
			username = args[0]
			if err := describeUserValidateUsername(username); err != nil {
				log.Fatalln(err)
			}
		} else {
			username = global.Prompt(promptui.Prompt{}, "Username:", describeUserValidateUsername)
		}
		var user db.User = db.GetUsers(username, 0)[0]
		printFormatted(username, []string{"Id", "Username"}, []interface{}{
			user.ID,
			user.Username,
		})
	},
}

func describeUserValidateUsername(username string) error {
	if !db.CheckUserExists(username) {
		return errors.New("user does not exist")
	}
	return nil
}

func init() {
	DescribeCmd.AddCommand(userCmd)
}
