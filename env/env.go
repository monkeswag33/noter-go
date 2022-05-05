package env

import (
	"github.com/monkeswag33/noter-go/types"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func SetupViper() error {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	viper.AutomaticEnv()
	viper.SetDefault("LOG_LEVEL", "warn")
	return nil
}

func ParseLogLevel() types.LogLevelParams {
	// Get default log level and gorm log level
	var logLevel = viper.GetString("LOG_LEVEL")
	var gormLogLevel string = viper.GetString("GORM_LOG_LEVEL")
	if len(gormLogLevel) == 0 {
		gormLogLevel = logLevel
	}
	return types.LogLevelParams{
		LogLevel:     logLevel,
		GormLogLevel: gormLogLevel,
	}
}

func SetLogLevel(params types.LogLevelParams) error {
	var ll logrus.Level
	ll, err := logrus.ParseLevel(params.LogLevel)
	if err != nil {
		logrus.Warnf("Unrecognized log level %q, using default value WARN", params.LogLevel)
		ll = logrus.WarnLevel
	}
	logrus.SetLevel(ll)
	logrus.Info("Set log level...")
	logrus.Debugf("Log level: %q", ll.String())
	return nil
}
