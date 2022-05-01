package db

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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

type DB struct {
	DB       *gorm.DB
	LogLevel string
}

func (db *DB) getLogLevel() logger.Interface {
	var loggerLevel logger.LogLevel
	logrus.Debugf("GORM logging level is: %q", db.LogLevel)
	switch strings.ToLower(db.LogLevel) {
	case "trace", "debug", "info":
		loggerLevel = logger.Info
	case "warn", "":
		loggerLevel = logger.Warn
	case "error":
		loggerLevel = logger.Error
	case "silent", "fatal": // Gorm log level silent, or logrus log level fatal
		loggerLevel = logger.Silent
	}
	logrus.Warnf("Unrecognized gorm log level %q, using default value WARN", db.LogLevel)
	loggerLevel = logger.Warn
	return logger.Default.LogMode(loggerLevel)
}

func (db *DB) Init() {
	if db.DB == nil {
		logrus.Info("Looking for POSTGRES_URI environment variable...")
		var uri string = viper.GetString("POSTGRES_URI")
		if len(uri) == 0 {
			logrus.Fatal("Could not find POSTGRES_URI environment variable. Please set it in a .env file or as an environment variable")
		}
		logrus.Info("Found POSTGRES_URI")
		logrus.Debugf("POSTGRES_URI is: %q", uri)
		var loggerLevel logger.Interface = db.getLogLevel()
		logrus.Info("Got GORM log level")

		logrus.Info("Connecting to database...")
		var err error
		db.DB, err = gorm.Open(postgres.Open(uri), &gorm.Config{
			Logger: loggerLevel,
		})
		logrus.Info("Connected to database")

		if err != nil {
			logrus.Fatal("Failed to connect to the database")
		}
	}

	logrus.Info("Running migrations...")
	db.DB.AutoMigrate(&User{})
	logrus.Trace("Migrated users...")
	db.DB.AutoMigrate(&Note{})
	logrus.Trace("Migrated notes...")
	logrus.Info("Finished migrations")
}

func (db *DB) Close() {
	var err error
	sqlDB, dbErr := db.DB.DB() // Get the actual sqlDB of the database
	if dbErr == nil {
		closeErr := sqlDB.Close()
		if err == nil {
			logrus.Info("Closed the database")
			return
		} else {
			err = closeErr
		}
	} else {
		err = dbErr
	}
	logrus.Error(err)
}
