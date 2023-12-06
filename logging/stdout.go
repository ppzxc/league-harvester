package logging

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type Level string

const (
	Panic Level = "panic"
	Fatal Level = "fatal"
	Error Level = "error"
	Warn  Level = "warn"
	Info  Level = "info"
	Debug Level = "debug"
	Trace Level = "trace"
)

type Option struct {
	UseJsonFormatter bool
	UseStdout        bool
	Level            Level
	LogFilename      string
	LogMaxSize       int
	LogMaxBackup     int
	LogMaxAge        int
	LogCompress      bool
}

func Setup(option Option) {
	if option.UseJsonFormatter {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{})
	}

	if option.UseStdout {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(&lumberjack.Logger{
			Filename:   option.LogFilename,
			MaxSize:    option.LogMaxBackup,
			MaxBackups: option.LogMaxBackup,
			MaxAge:     option.LogMaxAge,
			Compress:   option.LogCompress,
		})
	}

	switch option.Level {
	case Panic:
		log.SetLevel(log.PanicLevel)
	case Fatal:
		log.SetLevel(log.FatalLevel)
	case Error:
		log.SetLevel(log.ErrorLevel)
	case Warn:
		log.SetLevel(log.WarnLevel)
	case Info:
		log.SetLevel(log.InfoLevel)
	case Debug:
		log.SetLevel(log.DebugLevel)
	case Trace:
		log.SetLevel(log.TraceLevel)
	}
}

func LevelParser(level string) Level {
	switch level {
	case "panic":
		return Panic
	case "fatal":
		return Fatal
	case "error":
		return Error
	case "warn":
		return Warn
	case "info":
		return Info
	case "debug":
		return Debug
	case "trace":
		return Trace
	default:
		return Info
	}
}
