package ilog

// Level is the log level.
type Level int

// log level enum.
const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

var (
	levelBytes = map[Level][]byte{
		LevelDebug: []byte("Debug"),
		LevelInfo:  []byte("Info"),
		LevelWarn:  []byte("Warn"),
		LevelError: []byte("Error"),
		LevelFatal: []byte("Fatal"),
	}

	levelColor = map[Level]int{
		LevelDebug: 32, // green
		LevelInfo:  34, // blue
		LevelWarn:  33, // yellow
		LevelError: 31, // red
		LevelFatal: 35, // magenta
	}
)

// LogWriter defines writer's API.
type LogWriter interface {
	Init() error
	SetLevel(Level)
	GetLevel() Level
	Write(msg string, level Level) error
	Flush() error
	Close() error
}
