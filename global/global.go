package global

import (
	"log"

	"github.com/manifoldco/promptui"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Templates *promptui.PromptTemplates = &promptui.PromptTemplates{
	Prompt:  "{{ . }} ",
	Valid:   "{{ . | green }} ",
	Invalid: "{{ . | red }} ",
	Success: "{{ . | bold }} ",
}

func Prompt(config promptui.Prompt, label string, validator func(string) error) string {
	config.Label = label
	config.Templates = Templates
	config.Validate = validator
	result, err := config.Run()
	if err != nil {
		log.Fatalf("Error while getting user input: %v\n", err)
	}
	return result
}

func SetupViper() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.AutomaticEnv()
	viper.SetDefault("LOG_LEVEL", "warn")
}

func parseLogLevel() (string, string) {
	// Get default log level and gorm log level
	var logLevel = viper.GetString("LOG_LEVEL")
	var gormLogLevel string = viper.GetString("GORM_LOG_LEVEL")
	if len(gormLogLevel) == 0 {
		gormLogLevel = logLevel
	}
	return logLevel, gormLogLevel
}

func SetLogLevel() string {
	logLevel, gormLogLevel := parseLogLevel()
	ll, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.Fatalf("%q is not a valid log level", logLevel)
	}
	logrus.SetLevel(ll)
	logrus.Info("Set log level...")
	logrus.Debugf("Log level is: %q", logLevel)
	return gormLogLevel
}
