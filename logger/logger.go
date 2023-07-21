package logger

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

type FormatterHook struct {
	Writer    io.Writer
	LogLevels []log.Level
	Formatter log.Formatter
}

func (hook *FormatterHook) Fire(entry *log.Entry) error {
	line, err := hook.Formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write(line)
	return err
}

func (hook *FormatterHook) Levels() []log.Level {
	return hook.LogLevels
}

func InitLogger(f *os.File, logLevel string) {
	ll, err := log.ParseLevel(logLevel)
	if err != nil {
		panic(err)
	}
	log.SetOutput(io.Discard)
	log.SetLevel(ll)

	log.AddHook(&FormatterHook{
		Writer: os.Stderr,
		LogLevels: []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
			log.WarnLevel,
			log.InfoLevel,
		},
		Formatter: &log.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceColors:     true,
		},
	})
	log.AddHook(&FormatterHook{
		Writer: f,
		LogLevels: []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
			log.WarnLevel,
			log.InfoLevel,
			log.DebugLevel,
		},
		Formatter: &log.JSONFormatter{},
	})
}
