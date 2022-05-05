/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package new

import (
	"fmt"
	"regexp"

	"github.com/manifoldco/promptui"
	"github.com/monkeswag33/noter-go/argon2"
	"github.com/monkeswag33/noter-go/db"
	"github.com/monkeswag33/noter-go/errordef"
	"github.com/monkeswag33/noter-go/prompt"
	"github.com/monkeswag33/noter-go/types"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

const minEntropyBits float64 = 60

// new/userCmd represents the new/user command
var userCmd = &cobra.Command{
	Use:   "user <username>",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		database = db.Database
		password, _ := cmd.Flags().GetString("password")
		user, err := createUser(args, password)
		if err != nil {
			logrus.Fatal(err)
		}
		if err := insertUser(user); err != nil {
			logrus.Fatal(err)
		}
		fmt.Printf("Created user %s\n", user.Username)
	},
}

func createUser(args []string, password string) (*db.User, error) {
	var username string
	if len(args) == 1 {
		logrus.Debug("Recognized username as an argument, using it...")
		username = args[0]
		if err := newUserValidateUsername(username); err != nil {
			return nil, err
		}
	} else {
		logrus.Debug("Username was not given, prompting for it")
		username = prompt.Prompt(promptui.Prompt{}, "Username:", newUserValidateUsername)
	}
	logrus.Debug("Username passed validation")
	if len(password) == 0 {
		logrus.Debug("Password was not given as parameter, prompting for it")
		password = prompt.Prompt(promptui.Prompt{
			Mask: '*',
		}, "Password:", newUserValidatePassword)
	} else if err := newUserValidatePassword(password); err != nil {
		return nil, err
	}
	logrus.Debug("Password passed validation")
	hash, err := argon2.HashPass(password, &types.HashParams{
		Memory:      64 * 1024,
		Iterations:  4,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	})
	logrus.Debug("Password hashed...")
	if err != nil {
		return nil, err
	}
	var user db.User = db.User{
		Username: username,
		Password: hash,
	}
	logrus.Info("Created user")
	return &user, nil
}

func insertUser(user *db.User) error {
	if err := database.CreateUser(user); err != nil {
		return err
	}
	return nil
}

func newUserValidateUsername(username string) error {
	if len(username) < 3 {
		return errordef.ErrUsernameTooShort
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(username) {
		return errordef.ErrUsernameMustContainAlphaNumeric
	}
	exists, err := database.CheckUserExists(db.User{
		Username: username,
	})
	if err != nil {
		logrus.Fatal(err)
	}
	if exists {
		return errordef.ErrUserAlreadyExists
	}
	return nil
}

func newUserValidatePassword(password string) error {
	if len(password) < 8 {
		return errordef.ErrPasswordTooShort
	}
	if err := passwordvalidator.Validate(password, minEntropyBits); err != nil {
		return err
	}
	return nil
}

func init() {
	NewCmd.AddCommand(userCmd)
	userCmd.Flags().StringP("password", "p", "", "Password for new user")
}
