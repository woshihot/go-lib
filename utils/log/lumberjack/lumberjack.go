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
	logger *lumberjack.Logger
}

func Init() Logger {
	return Logger{&lumberjack.Logger{}}
}

func (l Logger) MaxSize(size int) Logger {
	l.logger.MaxSize = size
	return l
}

func (l Logger) MaxBackups(bu int) Logger {
	l.logger.MaxBackups = bu
	return l
}

func (l Logger) MaxAge(age int) Logger {
	l.logger.MaxAge = age
	return l
}

func (l Logger) LocalTime(localTime bool) Logger {
	l.logger.LocalTime = localTime
	return l
}

func (l Logger) Compress(compress bool) Logger {
	l.logger.Compress = compress
	return l
}

func (l Logger) FileName(fileName string) Logger {
	l.logger.Filename = fileName
	return l
}

func (l Logger) IoWriter() io.Writer {
	return l.logger
}
