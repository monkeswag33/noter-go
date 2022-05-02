package global

import (
	"os"
	"testing"

	"github.com/monkeswag33/noter-go/types"
	"github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"
)

var (
	logLevels [6]string = [...]string{
		"trace", "debug", "info",
		"warn", "error", "fatal",
	}

	gormLogLevels [4]string = [...]string{
		"info", "warn", "error", "silent",
	}
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
	for _, logLevel := range logLevels {
		for _, gormLogLevel := range gormLogLevels {
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
	for _, logLevel := range logLevels {
		var logLevelParams types.LogLevelParams = types.LogLevelParams{
			LogLevel: logLevel,
		}
		assert.NoError(t, SetLogLevel(logLevelParams))
		lvl, err := logrus.ParseLevel(logLevel)
		assert.NoError(t, err)
		assert.Equal(t, logrus.GetLevel(), lvl)
	}
}
