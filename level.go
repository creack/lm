package lm

// Level type.
type Level int

// LogLvel enum.
const (
	_ Level = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
	LevelFatal
)

func (l Level) String() string {
	switch l {
	case LevelFatal:
		return "FATAL"
	case LevelError:
		return "ERROR"
	case LevelWarning:
		return "WARNING"
	case LevelInfo:
		return "INFO"
	default:
		return "DEBUG"
	}
}
