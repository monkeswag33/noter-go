package db

import (
	"github.com/monkeswag33/noter-go/types"
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
	LogLevel types.LogLevelParams
}

func (db *DB) getLogLevel() logger.Interface {
	var loggerLevel logger.LogLevel
	if val, ok := types.GormLogLevels[db.LogLevel.GormLogLevel]; ok {
		loggerLevel = val
	} else {
		logrus.Warnf("Unrecognized gorm log level %q, using default value WARN", db.LogLevel.GormLogLevel)
		db.LogLevel.GormLogLevel = "warn"
		loggerLevel = logger.Warn
	}
	logrus.Debugf("GORM log level: %q", db.LogLevel.GormLogLevel)
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
