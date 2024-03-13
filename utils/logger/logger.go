package logger

import (
	"github.com/gookit/color"
	"log"
	"os"
)

var defaultLogger = log.New(os.Stderr, "", log.Lshortfile)

func Debugln(v ...any) {
	defaultLogger.Output(2, color.Green.Renderln(v))
}

// Debugf don't need to add `\n`. it's automatically added.
func Debugf(format string, v ...any) {
	defaultLogger.Output(2, color.Green.Sprintf(format, v))
}

func Infoln(v ...any) {
	defaultLogger.Output(2, color.Blue.Renderln(v))
}

// Infof don't need to add `\n`. it's automatically added.
func Infof(format string, v ...any) {
	defaultLogger.Output(2, color.Blue.Sprintf(format, v))
}

func Warnln(v ...any) {
	defaultLogger.Output(2, color.Magenta.Renderln(v))
}

// Warnf don't need to add `\n`. it's automatically added.
func Warnf(format string, v ...any) {
	defaultLogger.Output(2, color.Magenta.Sprintf(format, v))
}

func Errorln(v ...any) {
	defaultLogger.Output(2, color.Red.Renderln(v))
}

// Errorf don't need to add `\n`. it's automatically added.
func Errorf(format string, v ...any) {
	defaultLogger.Output(2, color.Red.Sprintf(format, v))
}
