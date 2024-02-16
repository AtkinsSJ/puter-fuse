package debug

import (
	"fmt"

	"github.com/HeyPuter/puter-fuse-go/services"
)

type ILogger interface {
	Log(format string, args ...interface{})
	Sub(crumbs []string) *Logger
	S(str string) *Logger
}

type Logger struct {
	Crumbs []string
}

func NewLogger(format string, args ...interface{}) *Logger {
	return &Logger{
		Crumbs: []string{fmt.Sprintf(format, args...)},
	}
}

func (l *Logger) Log(format string, args ...interface{}) {
	str := ""
	for _, crumb := range l.Crumbs {
		str += fmt.Sprintf("[%s] ", crumb)
	}

	// Wrap in blue
	str = fmt.Sprintf("\033[34;1m%s\033[0m", str)

	str += fmt.Sprintf(format, args...)

	fmt.Println(str)
}

func (l *Logger) Sub(crumbs []string) *Logger {
	return &Logger{
		Crumbs: append(l.Crumbs, crumbs...),
	}
}

func (l *Logger) S(str string) *Logger {
	return &Logger{
		Crumbs: append(l.Crumbs, str),
	}
}

type LogService struct {
	Logger *Logger

	services services.IServiceContainer
}

func (svc *LogService) Init(services services.IServiceContainer) {
	svc.Logger = &Logger{}
}

func (svc *LogService) Log(msg string) {
	svc.Logger.Log(msg)
}

func (svc *LogService) GetLogger(format string, args ...interface{}) *Logger {
	logger := &Logger{
		Crumbs: append(svc.Logger.Crumbs, fmt.Sprintf(format, args...)),
	}
	return logger
}
