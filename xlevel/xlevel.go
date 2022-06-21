package xlevel

import "strings"

// Level describes output levels for logs and regular prints.
type Level uint8

const (
	// Text describes regular print without log formatting.
	Text Level = 1 << iota
	// Trace describes some information for debug purposes but with lower priority
	// than LevelDebug like "http request with context token".
	Trace
	// Debug describes some debug information like "http request with params".
	Debug
	// Info describes regular information like "service running/stopped".
	Info
	// Warn describes not errors but something important which might become errors
	// like "unexpected http content type".
	Warn
	// Error describes regular not fatal errors in cases when operation failed
	// but service continue working like "http 500 can't process file".
	Error
	// Fatal describes errors for cases when there is no reason to continue work,
	// and the service must be stopped smoothly like "no source file to process".
	// BE AWARE: This level usually calls os.Exit().
	Fatal
	// Panic describes errors for cases when something extremely unexpected happened,
	// and the service must be panic like "nil pointer call".
	// BE AWARE: This level usually calls panic().
	Panic
)

// Has checks that this mask contains specific level.
func (mask Level) Has(level Level) bool {
	if mask == 0 || level == 0 {
		return mask == level
	}
	return level&mask == level
}

// String returns string representation of the mask.
// If mask contains multiple levels they will be separated with pipeline.
// It will looks like "error|fatal|panic".
// Empty string if level is zero.
func (mask Level) String() string {
	names := make([]string, 0, 8)
	levels := []Level{
		Text, Trace, Debug, Info,
		Warn, Error, Fatal, Panic,
	}
	for i := range levels {
		if mask.Has(levels[i]) {
			name := ""
			switch levels[i] {
			case Text:
				name = "text"
			case Trace:
				name = "trace"
			case Debug:
				name = "debug"
			case Info:
				name = "info"
			case Warn:
				name = "warn"
			case Error:
				name = "error"
			case Fatal:
				name = "fatal"
			case Panic:
				name = "panic"
			}
			names = append(names, name)
		}
	}
	return strings.Join(names, "|")
}

// Elevate accepts simple level and applies all levels upper.
//  LevelError.Elevate() == LevelError | LevelFatal | LevelPanic
func (mask Level) Elevate() Level {
	levels := []Level{
		Text, Trace, Debug, Info,
		Warn, Error, Fatal, Panic,
	}
	for i := range levels {
		if levels[i] > mask {
			mask |= levels[i]
		}
	}
	return mask
}

// Lower returns minimal level from current mask.
func (mask Level) Lower() Level {
	levels := []Level{
		Text, Trace, Debug, Info,
		Warn, Error, Fatal, Panic,
	}
	for i := range levels {
		if mask.Has(levels[i]) {
			return levels[i]
		}
	}
	return 0
}

// Higher returns maximal level from current mask.
func (mask Level) Higher() Level {
	levels := []Level{
		Panic, Fatal, Error, Warn,
		Info, Debug, Trace, Text,
	}
	for i := range levels {
		if mask.Has(levels[i]) {
			return levels[i]
		}
	}
	return 0
}

// Parse tries to parse level name into simple LevelMask.
// Case insensitive. Returns 0 if can't parse.
// Multiple levels can be presented like: "error|fatal|panic".
func Parse(name string) Level {
	level := Level(0)
	names := strings.Split(strings.ToLower(name), "|")
	for i := range names {
		switch strings.TrimSpace(names[i]) {
		case "text":
			level |= Text
		case "trace":
			level |= Trace
		case "debug":
			level |= Debug
		case "information", "info":
			level |= Info
		case "warning", "warn":
			level |= Warn
		case "error":
			level |= Error
		case "fatal":
			level |= Fatal
		case "panic":
			level |= Panic
		}
	}
	return level
}
