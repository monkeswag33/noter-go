package global

import (
	"os"
	"testing"

	"github.com/monkeswag33/noter-go/types"
	"github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"
)

func TestViperSetup(t *testing.T) {
	var filename string = ".env"
	file, err := os.Create(filename)
	assert.NoError(t, err)
	file.Close()
	assert.NoError(t, SetupViper())
	assert.NoError(t, os.Remove(filename))
}

func TestParseLogLevel(t *testing.T) {
	var (
		val    string
		exists bool
	)
	val, exists = os.LookupEnv("LOG_LEVEL")
	if exists {
		defer os.Setenv("LOG_LEVEL", val) // Reset LOG_LEVEL to original state
	}
	val, exists = os.LookupEnv("GORM_LOG_LEVEL")
	if exists {
		defer assert.NoError(t, os.Setenv("GORM_LOG_LEVEL", val))
	}
	for logLevel := range types.LogLevels {
		for gormLogLevel := range types.GormLogLevels {
			assert.NoError(t, os.Setenv("LOG_LEVEL", logLevel))
			assert.NoError(t, os.Setenv("GORM_LOG_LEVEL", gormLogLevel))
			assert.Equal(t, ParseLogLevel(), types.LogLevelParams{
				LogLevel:     logLevel,
				GormLogLevel: gormLogLevel,
			})
		}
	}
}

func TestSetLogLevel(t *testing.T) {
	for logLevel, lvl := range types.LogLevels {
		var logLevelParams types.LogLevelParams = types.LogLevelParams{
			LogLevel: logLevel,
		}
		assert.NoError(t, SetLogLevel(logLevelParams))
		assert.Equal(t, logrus.GetLevel(), lvl)
	}
}
