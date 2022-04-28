/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package new

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"regexp"

	"github.com/manifoldco/promptui"
	"github.com/monkeswag33/noter-go/db"
	"github.com/monkeswag33/noter-go/global"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	passwordvalidator "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/argon2"
)

const (
	memory         uint32  = 64 * 1024
	iterations     uint32  = 3
	parallelism    uint8   = 2
	saltLength     uint32  = 16
	keyLength      uint32  = 32
	minEntropyBits float64 = 60
)

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
		var username string
		if len(args) == 1 {
			logrus.Debug("Recognized username as an argument, using it...")
			username = args[0]
			if err := newUserValidateUsername(username); err != nil {
				logrus.Fatal(err)
			}
		} else {
			logrus.Debug("Username was not given, prompting for it")
			username = global.Prompt(promptui.Prompt{}, "Username:", newUserValidateUsername)
		}
		logrus.Debug("Username passed validation")
		password, _ := cmd.Flags().GetString("password")
		if len(password) == 0 {
			logrus.Debug("Password was not given as parameter, prompting for it")
			password = global.Prompt(promptui.Prompt{
				Mask: '*',
			}, "Password:", newUserValidatePassword)
		} else if err := newUserValidatePassword(password); err != nil {
			logrus.Fatal(err)
		}
		logrus.Debug("Password passed validation")
		hash, err := hashPass(password)
		logrus.Debug("Password hashed...")
		if err != nil {
			logrus.Fatal(err)
		}
		user, err := db.CreateUser(username, hash)
		logrus.Info("Inserted user")
		if err != nil {
			logrus.Fatal(err)
		}
		fmt.Printf("Created user %s\n", user.Username)
	},
}

func newUserValidateUsername(username string) error {
	if len(username) < 3 {
		return errors.New("username must be at least 4 characters")
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(username) {
		return errors.New("username can only contain alphanumeric characters")
	}
	if db.CheckUserExists(username) {
		return errors.New("user already exists")
	}
	return nil
}

func newUserValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
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

func hashPass(password string) (encodedHash string, err error) {
	salt, err := genSalt()
	if err != nil {
		return "", err
	}
	logrus.Tracef("Generated salt %d bytes long", saltLength)

	var hash []byte = argon2.IDKey([]byte(password),
		salt,
		iterations,
		memory,
		parallelism,
		keyLength)
	logrus.Trace("Hashed password using argon2 to bytes")
	var b64Salt string = base64.RawStdEncoding.EncodeToString(salt)
	var b64Hash string = base64.RawStdEncoding.EncodeToString(hash)
	logrus.Trace("Converted salt and hash to strings")
	encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, memory, iterations, parallelism, b64Salt, b64Hash)
	logrus.Trace("Generated string hash")
	return encodedHash, nil
}

func genSalt() ([]byte, error) {
	bytes := make([]byte, saltLength)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
