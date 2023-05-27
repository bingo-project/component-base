package console

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

func Info(msg string) {
	colorOut(msg, "green")
}

func Error(msg string) {
	colorOut(msg, "red")
}

func Warn(msg string) {
	colorOut(msg, "yellow")
}

func Exit(msg string) {
	Error(msg)
	os.Exit(1)
}

func ExitIf(err error) {
	if err != nil {
		Exit(err.Error())
	}
}

func colorOut(message, color string) {
	_, _ = fmt.Fprintln(os.Stdout, ansi.Color(message, color))
}
