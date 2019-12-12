package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

const (
	Ldate         = log.Ldate         // the date in the local device zone: 2009/01/23
	Ltime         = log.Ltime         // the device in the local device zone: 01:23:23
	Lmicroseconds = log.Lmicroseconds // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile     = log.Llongfile     // full file name and line number: /a/b/c/d.go:23
	Lshortfile    = log.Lshortfile    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC          = log.LUTC          // if Ldate or Ltime is set, use UTC rather than the local device zone
	LstdFlags     = log.LstdFlags     // initial values for the standard logger
)

type Logger struct {
	log.Logger
	level  LogLevel
	mu     sync.Mutex
	prefix string
}

func New(out io.Writer, level LogLevel, flag int) *Logger {
	return &Logger{Logger: *log.New(out, "", flag), level: level}
}

var std = New(os.Stderr, Verbose, log.LstdFlags)

func (l *Logger) setPrefix(p string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.prefix = p
}

func (l *Logger) setLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

func (l *Logger) isLeverContain(level LogLevel) bool {

	return l.level <= level
}

func SetLevel(level LogLevel) {
	std.setLevel(level)
}

func D(v ...interface{}) {
	if std.isLeverContain(Debug) {
		prefix(d)
		write(fmt.Sprint(v...))
	}
}

func I(v ...interface{}) {
	if std.isLeverContain(Info) {
		prefix(i)
		write(fmt.Sprint(v...))
	}
}

func W(v ...interface{}) {
	if std.isLeverContain(Warn) {
		prefix(w)
		write(fmt.Sprint(v...))
	}
}

func E(v ...interface{}) {
	if std.isLeverContain(Error) {
		prefix(e)
		write(fmt.Sprint(v...))
	}
}

func Df(format string, v ...interface{}) {
	if std.isLeverContain(Debug) {
		prefix(d)
		write(fmt.Sprintf(format, v...))
	}
}

func If(format string, v ...interface{}) {
	if std.isLeverContain(Info) {
		prefix(i)
		write(fmt.Sprintf(format, v...))
	}
}

func Wf(format string, v ...interface{}) {
	if std.isLeverContain(Warn) {
		prefix(w)
		write(fmt.Sprintf(format, v...))
	}
}

func Ef(format string, v ...interface{}) {
	if std.isLeverContain(Error) {
		prefix(e)
		write(fmt.Sprintf(format, v...))
	}
}

func EF(tag, format string, v ...interface{}) {
	if std.isLeverContain(Error) {
		prefix(e)
		write(fmt.Sprintf(tag+" "+format, v...))
	}
}
func DF(tag, format string, v ...interface{}) {
	if std.isLeverContain(Debug) {
		prefix(d)
		write(fmt.Sprintf(tag+" "+format, v...))
	}
}
func IF(tag, format string, v ...interface{}) {
	if std.isLeverContain(Info) {
		prefix(i)
		write(fmt.Sprintf(tag+" "+format, v...))
	}
}
func WF(tag, format string, v ...interface{}) {
	if std.isLeverContain(Warn) {
		prefix(w)
		write(fmt.Sprintf(tag+" "+format, v...))
	}
}

func Dln(v ...interface{}) {
	if std.isLeverContain(Debug) {
		prefix(d)
		write(fmt.Sprintln(v...))
	}
}

func Iln(v ...interface{}) {
	if std.isLeverContain(Info) {
		prefix(i)
		write(fmt.Sprintln(v...))
	}
}

func Wln(v ...interface{}) {
	if std.isLeverContain(Warn) {
		prefix(w)
		write(fmt.Sprintln(v...))
	}
}

func Eln(v ...interface{}) {
	if std.isLeverContain(Error) {
		prefix(e)
		write(fmt.Sprintln(v...))
	}
}

// Panic is equivalent to Print() followed by a call to panic().
func Panic(v ...interface{}) {
	w := fmt.Sprint(v...)
	prefix(e)
	write(w)
	panic(w)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Panicf(format string, v ...interface{}) {
	w := fmt.Sprintf(format, v...)
	prefix(e)
	write(w)
	panic(w)
}

// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(v ...interface{}) {
	w := fmt.Sprintln(v...)
	prefix(e)
	write(w)
	panic(w)
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	prefix(e)
	write(fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	prefix(e)
	write(fmt.Sprintf(format, v...))
	os.Exit(1)

}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(v ...interface{}) {
	prefix(e)
	write(fmt.Sprintln(v...))
	os.Exit(1)
}

func write(s string) {
	std.Output(3, s)
}

func prefix(s string) {
	std.Logger.SetPrefix(std.prefix + s)
}

// SetOutput sets the output destination for the standard logger.
func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

// Flags returns the output flags for the standard logger.
func Flags() int {
	return std.Flags()
}

// SetFlags sets the output flags for the standard logger.
func SetFlags(flag int) {
	std.SetFlags(flag)
}

// Prefix returns the output prefix for the standard logger.
func Prefix() string {
	return std.Prefix()
}

// SetPrefix sets the output prefix for the standard logger.
func SetPrefix(prefix string) {

	std.setPrefix(prefix)
}
