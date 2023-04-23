package log

import "github.com/blackmarllboro/create-project-struct/internal/pkg/log/interfaces"

type Log struct {
	interfaces.Logger
}

func NewLogger(logger interfaces.Logger) *Log {
	return &Log{Logger: logger}
}
