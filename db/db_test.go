package db

import (
	"os"
	"testing"

	"github.com/monkeswag33/noter-go/types"
	"github.com/stretchr/testify/assert"
)

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
	_, exists := os.LookupEnv("POSTGRES_URI")
	t.Fatal(exists)
	// global.SetupViper()
	// var database DB = DB{
	// 	LogLevel: types.LogLevelParams{
	// 		LogLevel:     "warn",
	// 		GormLogLevel: "warn",
	// 	},
	// }
	// assert.NoError(t, database.Init())
}

// func TestSetupDB(t *testing.T) {
// 	global.SetupViper()
// 	assert.NoError(t, SetupDB(types.LogLevelParams{
// 		LogLevel:     "warn",
// 		GormLogLevel: "warn",
// 	}))
// 	assert.IsType(t, &DB{}, Database)
// }
