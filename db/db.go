package db

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

type Note struct {
	ID     int
	Name   string `gorm:"unique;not null"`
	Body   string `gorm:"not null"`
	UserID int    `gorm:"not null"`
	User   User   `gorm:"constraint:OnDelete:CASCADE;"`
}

type User struct {
	ID       int
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func getLogLevel(logLevel string) logger.Interface {
	var loggerLevel logger.LogLevel
	logrus.Debugf("GORM logging level is: %q", logLevel)
	switch strings.ToLower(logLevel) {
	case "trace", "debug", "info":
		loggerLevel = logger.Info
	case "warn":
		loggerLevel = logger.Warn
	case "error":
		loggerLevel = logger.Error
	case "silent", "fatal": // Gorm log level silent, or logrus log level fatal
		loggerLevel = logger.Silent
	default:
		logrus.Warnf("Unrecognized gorm log level %q, using default value WARN", logLevel)
		loggerLevel = logger.Warn
	}
	return logger.Default.LogMode(loggerLevel) // This will never be called, just here to stop go from complaining about return types
}

func InitDB(logLevel string) {
	logrus.Info("Looking for POSTGRES_URI environment variable...")
	var uri string = viper.GetString("POSTGRES_URI")
	if len(uri) == 0 {
		logrus.Fatal("Could not find POSTGRES_URI environment variable. Please set it in a .env file or as an environment variable")
	}
	logrus.Info("Found POSTGRES_URI")
	logrus.Debugf("POSTGRES_URI is: %q", uri)

	var loggerLevel logger.Interface = getLogLevel(logLevel)
	logrus.Info("Got GORM log level")

	logrus.Info("Connecting to database...")
	var err error
	db, err = gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: loggerLevel,
	})
	logrus.Info("Connected to database")

	if err != nil {
		logrus.Fatal("Failed to connect to the database")
	}

	logrus.Info("Running migrations...")
	db.AutoMigrate(&User{})
	logrus.Trace("Migrated users...")
	db.AutoMigrate(&Note{})
	logrus.Trace("Migrated notes...")
	logrus.Info("Finished migrations")
}
