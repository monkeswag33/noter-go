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

var userCmd = &cobra.Command{
	Use:   "user <username>",
	Short: "Delete a user",
	Long:  "Delete a specific user by username as the first",
	Run: func(cmd *cobra.Command, args []string) {
		var username string
		if len(args) == 1 {
			logrus.Debug("Username recognized as argument, using it")
			username = args[0]
			if err := deleteUserValidateUsername(username); err != nil {
				logrus.Fatal(err)
			}
		} else {
			logrus.Debug("Username not found, prompting")
			username = global.Prompt(promptui.Prompt{}, "Username:", deleteUserValidateUsername)
		}
		logrus.Debug("Username passed validation")
		var deletedUser db.User = db.DeleteUser(username)
		logrus.Info("Deleted user")
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
