package log

type Log struct {
	Logger
}

func NewLogger(logger Logger) *Log {
	return &Log{Logger: logger}
}
