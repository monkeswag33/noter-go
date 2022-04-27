package main

import (
	"log"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

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
