package db

import (
	"github.com/monkeswag33/noter-go/errordef"
	"github.com/monkeswag33/noter-go/types"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
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

var Database *DB

func SetupDB(logLevel types.LogLevelParams) error {
	Database = &DB{
		LogLevel: logLevel,
	}
	if err := Database.Init(); err != nil {
		return err
	}
	return nil
}

func ShutdownDB() error {
	if err := Database.Close(); err != nil {
		return err
	}
	Database = nil
	return nil
}

func InitTesterDB() (*DB, error) {
	// Specify location to be in-memory
	// Cache is not shared, so each connection will be a seperate database
	var location string = ":memory:"
	gormDB, err := gorm.Open(sqlite.Open(location), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	var database DB = DB{
		LogLevel: types.LogLevelParams{
			LogLevel:     "warn",
			GormLogLevel: "info",
		},
		DB: gormDB,
	}
	if err := database.Init(); err != nil {
		return nil, err
	}
	return &database, nil
}

func (db *DB) getLogLevel() (loggerLevel logger.LogLevel) {
	if val, ok := types.GormLogLevels[db.LogLevel.GormLogLevel]; ok {
		loggerLevel = val
	} else {
		logrus.Warnf("Unrecognized gorm log level %q, using default value WARN", db.LogLevel.GormLogLevel)
		db.LogLevel.GormLogLevel = "warn"
		loggerLevel = logger.Warn
	}
	logrus.Debugf("GORM log level: %q", db.LogLevel.GormLogLevel)
	return loggerLevel
}

func (db *DB) Init() error {
	if db.DB == nil {
		logrus.Info("Looking for POSTGRES_URI environment variable...")
		var uri string = viper.GetString("POSTGRES_URI")
		if len(uri) == 0 {
			return errordef.ErrCouldNotFindPostgresURI
		}
		logrus.Info("Found POSTGRES_URI")
		logrus.Debugf("POSTGRES_URI is: %q", uri)
		var loggerLevel logger.Interface = logger.Default.LogMode(db.getLogLevel())
		logrus.Info("Got GORM log level")

		logrus.Info("Connecting to database...")
		var err error
		db.DB, err = gorm.Open(postgres.Open(uri), &gorm.Config{
			Logger: loggerLevel,
		})
		logrus.Info("Connected to database")

		if err != nil {
			return errordef.ErrFailedToConnect
		}
	}

	logrus.Info("Running migrations...")
	db.DB.AutoMigrate(&User{})
	logrus.Trace("Migrated users...")
	db.DB.AutoMigrate(&Note{})
	logrus.Trace("Migrated notes...")
	logrus.Info("Finished migrations")
	return nil
}

func (db *DB) Close() error {
	var err error
	sqlDB, dbErr := db.DB.DB() // Get the actual sqlDB of the database
	if dbErr == nil {
		closeErr := sqlDB.Close()
		if err == nil {
			logrus.Info("Closed the database")
			return nil
		} else {
			err = closeErr
		}
	} else {
		err = dbErr
	}
	return err
}
