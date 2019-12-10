package log

type LogLevel int

const (
	Verbose LogLevel = iota
	Debug
	Info
	Warn
	Error
)

const (
	verbose = "VERBOSE"
	debug   = "DEBUG"
	info    = "INFO"
	warn    = "WARN"
	error   = "ERROR"
)

const (
	d = "[" + debug + "] "
	i = "[" + info + "] "
	w = "[" + warn + "] "
	e = "[" + error + "] "
)

func ParseLevel(l string) LogLevel {
	level := Verbose
	switch l {
	case verbose:
		level = Verbose
	case debug:
		level = Debug
	case info:
		level = Info
	case warn:
		level = Warn
	case error:
		level = Error
	}
	return level
}

func (l *LogLevel) String() string {
	s := verbose
	switch *l {
	case Verbose:
		s = verbose
	case Debug:
		s = debug
	case Info:
		s = info
	case Warn:
		s = warn
	case Error:
		s = error
	}
	return s
}
