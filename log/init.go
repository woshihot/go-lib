package log

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
)

var LogOutput io.Writer

func Init(path, sysName string) {
	lumberjackLogger := &lumberjack.Logger{
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     3,
		LocalTime:  true,
		Compress:   true,
	}
	lumberjackLogger.Filename = path

	LogOutput = io.MultiWriter(os.Stdout, lumberjackLogger)
	log.SetPrefix("[" + sysName + "] ")
	log.SetOutput(LogOutput)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
