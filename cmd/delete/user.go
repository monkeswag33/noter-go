package delete

import (
	"errors"
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/monkeswag33/noter-go/db"
	"github.com/monkeswag33/noter-go/global"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user <username>",
	Short: "Delete a user",
	Long:  "Delete a specific user by username as the first",
	Run: func(cmd *cobra.Command, args []string) {
		var username string
		if len(args) == 1 {
			username = args[0]
			if err := deleteUserValidateUsername(username); err != nil {
				log.Fatalln(err)
			}
		} else {
			username = global.Prompt(promptui.Prompt{}, "Username:", deleteUserValidateUsername)
		}
		var deletedUser db.User = db.DeleteUser(username)
		fmt.Printf("Deleted user %s\n", deletedUser.Username)
	},
}

func deleteUserValidateUsername(username string) error {
	if !db.CheckUserExists(username) {
		return errors.New("user does not exist")
	}
	return nil
}

func init() {
	DeleteCmd.AddCommand(userCmd)
}
