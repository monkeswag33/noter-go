package db

import (
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/monkeswag33/noter-go/global"
	"github.com/monkeswag33/noter-go/types"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func changeDir(location string) string {
	_, current, _, _ := runtime.Caller(0)
	down := path.Join(path.Dir(current), location)
	if err := os.Chdir(down); err != nil {
		logrus.Fatal(err)
	}
	current, _ = os.Getwd()
	return current
}

func TestMain(m *testing.M) {
	changeDir("..")
	code := m.Run()
	os.Exit(code)
}

func TestGetLogLevel(t *testing.T) {
	var database DB
	for logLevel, lvl := range types.GormLogLevels {
		database.LogLevel.GormLogLevel = logLevel
		assert.Equal(t, lvl, database.getLogLevel())
	}
}

func TestInitTesterDB(t *testing.T) {
	database, err := InitTesterDB()
	assert.NoError(t, err)
	assert.IsType(t, &DB{}, database)
}

func TestInit(t *testing.T) {
	global.SetupViper()
	var database DB = DB{
		LogLevel: types.LogLevelParams{
			GormLogLevel: "warn",
		},
	}
	assert.NoError(t, database.Init())
	_, err := database.GetUsers(User{})
	assert.NoError(t, err)
}

func TestSetupDB(t *testing.T) {
	global.SetupViper()
	assert.NoError(t, SetupDB(types.LogLevelParams{
		GormLogLevel: "warn",
	}))
	assert.IsType(t, &DB{}, Database)
	_, err := Database.GetUsers(User{})
	assert.NoError(t, err)
}

func TestClose(t *testing.T) {
	SetupDB(types.LogLevelParams{
		GormLogLevel: "warn",
	})
	assert.NoError(t, Database.Close())
	sqlDB, err := Database.DB.DB()
	assert.NoError(t, err)
	assert.Error(t, sqlDB.Ping()) // Ping should error out if connection was closed
}

func TestShutdownDB(t *testing.T) {
	SetupDB(
		types.LogLevelParams{
			GormLogLevel: "warn",
		},
	)
	assert.NoError(t, ShutdownDB())
	if Database != nil {
		t.Fatal("Database is not null")
	}
}
