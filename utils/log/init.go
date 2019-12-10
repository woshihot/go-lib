package log

import (
	"io"
	"os"
)

var LogOutput io.Writer

// set sysName-prefix for log
// set flag for log
// set multiWriter for log
func Init(sysName string, flags int, writers ...io.Writer) {
	writers = append(writers, os.Stdout)
	LogOutput = io.MultiWriter(writers...)
	SetPrefix("[" + sysName + "] ")
	SetOutput(LogOutput)
	SetFlags(flags)
}
