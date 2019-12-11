package lumberjack

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
)

func IoWriter(path string, maxSize, maxBackups, maxAge int, localTime, compress bool) io.Writer {
	return &lumberjack.Logger{
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		LocalTime:  localTime,
		Compress:   compress,
		Filename:   path,
	}
}

type Logger struct {
	*lumberjack.Logger
}

func (l Logger) Init() Logger {
	return Logger{&lumberjack.Logger{}}
}

func (l Logger) MaxSize(size int) Logger {
	l.Logger.MaxSize = size
	return l
}

func (l Logger) MaxBackups(bu int) Logger {
	l.Logger.MaxBackups = bu
	return l
}

func (l Logger) MaxAge(age int) Logger {
	l.Logger.MaxAge = age
	return l
}

func (l Logger) LocalTime(localTime bool) Logger {
	l.Logger.LocalTime = localTime
	return l
}

func (l Logger) Compress(compress bool) Logger {
	l.Logger.Compress = compress
	return l
}

func (l Logger) FileName(fileName string) Logger {
	l.Logger.Filename = fileName
	return l
}

func (l Logger) IoWriter() io.Writer {
	return l.Logger
}
