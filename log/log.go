package log

import (
	"github.com/dawsonalex/todo-server/build"
	"github.com/sirupsen/logrus"
)

func New() *Logger {
	return &Logger{
		Logger: logrus.New(),
	}
}

type Logger struct {
	*logrus.Logger
}

func (l *Logger) WithBuildInfo() *Entry {
	b := build.Info()
	return l.WithFields(Fields{
		"version":     b.Version.Sprint(),
		"commit":      b.Commit,
		"branch":      b.Branch,
		"host":        b.Host,
		"environment": b.Environment,
	})
}
