/*
	The goal of the common package is to provide simple functions that will be used by other packages

Here we'll use Zerolog for Golang
It's a package that serves to write logs to some output (file, console,....)
Apparently it's a really fast json log writer
It provides easier reading for human and an easy pattern for automatic treatments
*/
package logger

import (
	"starter-pack-api/internal/config"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	logger zerolog.Logger
}

func OpenLogger(confLogs config.Logs) Logger {
	//Code for rolling logs
	lumberjackLogger := &lumberjack.Logger{
		Filename:   confLogs.Path,
		MaxSize:    1,
		MaxBackups: 3,
		MaxAge:     7,
		Compress:   true,
	}
	//Configuration of logs
	zerolog.SetGlobalLevel(zerolog.Level(confLogs.Level))
	return Logger{logger: zerolog.New(lumberjackLogger).With().Timestamp().Logger()}
}

func (l Logger) PanicLevel(message string, additionnalFields ...map[string]string) {
	l.logger.Panic().Fields(additionnalFields).Msg(message)
}

func (l Logger) FatalLevel(message string, additionnalFields ...map[string]string) {
	l.logger.Fatal().Fields(additionnalFields).Msg(message)
}

func (l Logger) ErrorLevel(message string, err error, additionnalFields ...map[string]string) {
	l.logger.Error().Fields(additionnalFields).Err(err).Msg(message)
}

func (l Logger) WarnLevel(message string, additionnalFields ...map[string]string) {
	l.logger.Warn().Fields(additionnalFields).Msg(message)
}

func (l Logger) InfoLevel(message string, additionnalFields ...map[string]interface{}) {
	l.logger.Info().Fields(additionnalFields).Msg(message)
}

func (l Logger) DebugLevel(message string, additionnalFields ...map[string]string) {
	l.logger.Debug().Fields(additionnalFields).Msg(message)
}

func (l Logger) TraceLevel(message string, additionnalFields ...map[string]string) {
	l.logger.Trace().Fields(additionnalFields).Msg(message)
}
