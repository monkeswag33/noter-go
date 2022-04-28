package describe

import (
	"errors"

	"github.com/manifoldco/promptui"
	"github.com/monkeswag33/noter-go/db"
	"github.com/monkeswag33/noter-go/global"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user <username>",
	Short: "Describe a specific user",
	Long:  "Get all the information about a specific user",
	Run: func(cmd *cobra.Command, args []string) {
		var username string
		if len(args) == 1 {
			logrus.Debug("Username recognized as argument, using it")
			username = args[0]
			if err := describeUserValidateUsername(username); err != nil {
				logrus.Fatal(err)
			}
		} else {
			logrus.Debug("Username not found, prompting")
			username = global.Prompt(promptui.Prompt{}, "Username:", describeUserValidateUsername)
		}
		logrus.Debug("Username passed validation")
		var user db.User = db.GetUsers(username, 0)[0]
		logrus.Debug("Received user")
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
