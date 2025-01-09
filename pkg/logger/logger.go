package logger

import (
	"fmt"
	"github.com/gookit/color"
	"log"
	"os"
	"strings"
)

func Debug(v ...any) {
	getLogger().Print(_DEBUG, fmt.Sprintln(v...))
}

// Debugf don't need to add `\n`. it's automatically added.
func Debugf(format string, v ...any) {
	getLogger().Print(_DEBUG, fmt.Sprintf(format, v...))
}

func Info(v ...any) {
	getLogger().Print(_INFO, fmt.Sprintln(v...))
}

// Infof don't need to add `\n`. it's automatically added.
func Infof(format string, v ...any) {
	getLogger().Print(_INFO, fmt.Sprintf(format, v...))
}

func Warn(v ...any) {
	getLogger().Print(_WARN, fmt.Sprintln(v...))
}

// Warnf don't need to add `\n`. it's automatically added.
func Warnf(format string, v ...any) {
	getLogger().Print(_WARN, fmt.Sprintf(format, v...))
}

func Error(v ...any) {
	getLogger().Print(_ERROR, fmt.Sprintln(v...))
}

// Errorf don't need to add `\n`. it's automatically added.
func Errorf(format string, v ...any) {
	getLogger().Print(_ERROR, fmt.Sprintf(format, v...))
}

type LogLevel uint8

const (
	_DEBUG LogLevel = iota
	_INFO
	_WARN
	_ERROR
)

type Logger interface {
	Print(logLevel LogLevel, v string)
}

func getLogger() Logger {
	return defaultLogger
}

type SimpleLogger struct {
	*log.Logger
	LogColor []color.Color
}

var defaultLogger = &SimpleLogger{
	log.New(os.Stderr, "", log.Lshortfile),
	[]color.Color{color.Green, color.Blue, color.Magenta, color.Red},
}

func (l *SimpleLogger) Print(logLevel LogLevel, v string) {
	_ = l.Output(3, Colored(l.LogColor[logLevel], v))
}

func Colored(renderingColor color.Color, v string) string {
	lines := strings.Split(v, "\n")
	for i, line := range lines {
		lines[i] = renderingColor.Sprintf(line)
	}
	return strings.Join(lines, "\n")
}
