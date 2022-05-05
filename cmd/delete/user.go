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

var userCmd = &cobra.Command{
	Use:   "user <username>",
	Short: "Delete a user",
	Long:  "Delete a specific user by username as the first",
	Run: func(cmd *cobra.Command, args []string) {
		database = db.Database
		var username string
		if len(args) == 1 {
			logrus.Debug("Username recognized as argument, using it")
			username = args[0]
			if err := deleteUserValidateUsername(username); err != nil {
				logrus.Fatal(err)
			}
		} else {
			logrus.Debug("Username not found, prompting")
			username = prompt.Prompt(promptui.Prompt{
				Label: "Username:",
			}, deleteUserValidateUsername)
		}
		logrus.Debug("Username passed validation")
		if err := database.DeleteUser(db.User{
			Username: username,
		}); err != nil {
			logrus.Fatal(err)
		}
		logrus.Info("Deleted user")
		fmt.Printf("Deleted user %s\n", username)
	},
}

func deleteUserValidateUsername(username string) error {
	exists, err := database.CheckUserExists(db.User{
		Username: username,
	})
	if err != nil {
		logrus.Fatal(err)
	}
	if !exists {
		return errordef.ErrUserDoesntExist
	}
	return nil
}

func init() {
	DeleteCmd.AddCommand(userCmd)
}
