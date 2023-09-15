package file

import (
	"os"

	"github.com/go-kratos/kratos/v2/log"
)

type fileLogger struct {
	f   *os.File
	log log.Logger
}

func NewFileLogger(f *os.File) log.Logger {
	return &fileLogger{
		f:   f,
		log: log.NewStdLogger(f),
	}
}

func (l *fileLogger) Log(level log.Level, keyvals ...interface{}) error {
	return l.log.Log(level, keyvals...)
}

func (l *fileLogger) Close() error {
	return nil
}
