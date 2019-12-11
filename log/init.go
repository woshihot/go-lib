package log

import (
	"github.com/woshihot/go-lib"
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
	go_lib.SetPrefix("[" + sysName + "] ")
	go_lib.SetOutput(LogOutput)
	go_lib.SetFlags(flags)
}
