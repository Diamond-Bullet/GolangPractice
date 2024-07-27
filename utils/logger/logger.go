package logger

import (
	"fmt"
	"github.com/gookit/color"
	"log"
	"os"
	"strings"
)

var defaultLogger = &Logger{
	log.New(os.Stderr, "", log.Lshortfile),
	[]color.Color{color.Green, color.Blue, color.Magenta, color.Red},
}

type LogLevel uint8

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
)

type Logger struct {
	*log.Logger
	LogColor []color.Color
}

func (l *Logger) Print(logLevel LogLevel, v string) {
	_ = l.Output(3, Colored(l.LogColor[logLevel], v))
}

func Debugln(v ...any) {
	defaultLogger.Print(Debug, fmt.Sprint(v))
}

// Debugf don't need to add `\n`. it's automatically added.
func Debugf(format string, v ...any) {
	defaultLogger.Print(Debug, fmt.Sprintf(format, v))
}

func Infoln(v ...any) {
	defaultLogger.Print(Info, fmt.Sprint(v))
}

// Infof don't need to add `\n`. it's automatically added.
func Infof(format string, v ...any) {
	defaultLogger.Print(Info, fmt.Sprintf(format, v))
}

func Warnln(v ...any) {
	defaultLogger.Print(Warn, fmt.Sprint(v))
}

// Warnf don't need to add `\n`. it's automatically added.
func Warnf(format string, v ...any) {
	defaultLogger.Print(Warn, fmt.Sprintf(format, v))
}

func Errorln(v ...any) {
	defaultLogger.Print(Error, fmt.Sprint(v))
}

// Errorf don't need to add `\n`. it's automatically added.
func Errorf(format string, v ...any) {
	defaultLogger.Print(Error, fmt.Sprintf(format, v))
}

func Colored(renderingColor color.Color, v string) string {
	lines := strings.Split(v, "\n")
	for i, line := range lines {
		lines[i] = renderingColor.Sprintf(line)
	}
	return strings.Join(lines, "\n")
}
