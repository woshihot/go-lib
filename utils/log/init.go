package log

import (
	"io"
	"log"
	"os"
)

var LogOutput io.Writer

func Init(sysName string, writers ...io.Writer) {
	writers = append(writers, os.Stdout)
	LogOutput = io.MultiWriter(writers...)
	log.SetPrefix("[" + sysName + "] ")
	log.SetOutput(LogOutput)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
