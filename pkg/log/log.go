package log

import (
	"changeme/internal/config"
	"errors"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func Init() (err error) {
	conf := config.Conf
	var out io.Writer
	if conf.Log.FilePath != "" {
		out = &lumberjack.Logger{
			Filename:   conf.Log.FilePath,
			MaxSize:    500, // megabytes
			MaxBackups: 3000,
			MaxAge:     3000, // days
			Compress:   true, // disabled by default
		}
	} else {
		out = os.Stdout
	}
	logrus.SetOutput(out)

	switch conf.App.Level {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		return errors.New("App level not in debug,info,warn,error")
	}

	return
}
