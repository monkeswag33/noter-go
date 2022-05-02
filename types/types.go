package types

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
)

type LogLevelParams struct {
	LogLevel     string
	GormLogLevel string
}

type HashParams struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

var (
	LogLevels map[string]logrus.Level = map[string]logrus.Level{
		"trace": logrus.TraceLevel,
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
	}
	GormLogLevels map[string]logger.LogLevel = map[string]logger.LogLevel{
		"trace":  logger.Info,
		"debug":  logger.Info,
		"info":   logger.Info,
		"warn":   logger.Warn,
		"error":  logger.Error,
		"silent": logger.Silent,
		"fatal":  logger.Silent,
	}
)
