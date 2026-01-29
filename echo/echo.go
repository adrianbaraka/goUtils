package echo

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/term"
)


// Global vars
const escape string = "\x1B["
const reset string = escape + "0m"

type colour string
type logLevel int

type logger struct {
	level    logLevel
	useColor bool
	out      io.Writer
}

// Returns a new logger with log level functionality.
func NewLogger(l logLevel, out io.Writer) *logger {
	return &logger{
		level:    l,
		useColor: shouldColor(out),
		out:      out,
	}
}

// colors
const (
	Red     colour = colour(escape + "38;5;196m")
	Green   colour = colour(escape + "92m")
	Yellow  colour = colour(escape + "93m")
	Blue    colour = colour(escape + "38;5;27m")
	Magenta colour = colour(escape + "95m")
	Cyan    colour = colour(escape + "96m")
)

const (
	Fatal logLevel = iota
	Error
	Warn
	Info
	Debug
	Trace
)

func shouldColor(w io.Writer) bool {
	if os.Getenv("FORCE_COLOR") != "" {
		//force-color.org
		// also for something like --color always
		return true
	}
	if os.Getenv("NO_COLOR") != "" {
		// no-color.org
		return false
	}
	// Check if the writer is a terminal
	if f, ok := w.(*os.File); ok {
		return term.IsTerminal(int(f.Fd()))
	}
	return false
}

// these adhere to the log level
// --------------------------------------------------------------------------------------------------

// Main function that handles the logic
func (l *logger) log(color colour, level logLevel, w io.Writer, format string, a ...any) (int, error) {
	if level > l.level {
		return 0, nil
	}
	// TODO add time
	if l.useColor {
		fmt.Fprintf(w, "%v", color) // color
		fmt.Fprintf(w, format, a...)
		return fmt.Fprintf(w, "%v", reset) // reset
	}
	return fmt.Fprintf(w, format, a...)
}

// Fechof formats and writes colored output to the provided writer with log level filtering. Overrides the initialized writer
func (l *logger) Fechof(color colour, level logLevel, w io.Writer, format string, a ...any) (int, error) {
	return l.log(color, level, w, format, a...)
}

// Fecholn writes colored output with a newline to the provided writer with log level filtering. Overrides the initialized writer
func (l *logger) Fecholn(color colour, level logLevel, w io.Writer, a ...any) (int, error) {
	msg := fmt.Sprintln(a...)
	return l.log(color, level, w, "%v", msg)
}

// Echoln writes colored output with a newline to configured writer with log level filtering.
func (l *logger) Echoln(color colour, level logLevel, a ...any) (int, error) {
	msg := fmt.Sprintln(a...)
	return l.log(color, level, l.out, "%v", msg)
}

// Echof formats and writes colored output to configured writer with log level filtering.
func (l *logger) Echof(color colour, level logLevel, format string, a ...any) (int, error) {
	return l.log(color, level, l.out, format, a...)
}

// some convenience methods

// Success message [Info]
func (l *logger) Success(a ...any) {
	msg := fmt.Sprintln(a...)
	l.Echof(Green, Info, "%v", msg)
}

// Debug message [Debug]
func (l *logger) Debug(a ...any) {
	msg := fmt.Sprintln(a...)
	l.Echof(Cyan, Debug, "%v", msg)
}

// error message [Error]
func (l *logger) Error(a ...any) {
	msg := fmt.Sprintln(a...)
	l.Echof(Red, Error, "%v", msg)
}

// --------------------------------------------------------------------------------------------------

// Just generic no logger is needed
// main function for the generic ones
func log(color colour, w io.Writer, format string, a ...any) (int, error) {
	l := NewLogger(Info, w)
	return l.log(color, Info, w, format, a...)
}

// Fechof formats and writes colored output to the provided writer.
func Fechof(color colour, w io.Writer, format string, a ...any) (int, error) {
	return log(color, w, format, a...)
}

// Fecholn writes colored output with a newline to the provided writer.
func Fecholn(color colour, w io.Writer, a ...any) (int, error) {
	msg := fmt.Sprintln(a...)
	return log(color, w, "%v", msg)
}

// Echoln writes colored output with a newline to stdout.
func Echoln(color colour, a ...any) (int, error) {
	msg := fmt.Sprintln(a...)
	return log(color, os.Stdout, "%v", msg)
}

// Echof formats and writes colored output to stdout.
func Echof(color colour, format string, a ...any) (int, error) {
	return log(color, os.Stdout, format, a...)
}
