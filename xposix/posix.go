package xposix

// Mode describes how must POSIX colors be processed for output writer.
type Mode uint8

const (
	// Auto tries to use POSIX colors and formats where it possible.
	// This option will remove format if NO_COLOR env var exists
	// or CLICOLOR=0 or if output target isn't *io.File which
	// supports POSIX.
	Auto Mode = iota
	// Forced pass POSIX as is.
	// Shouldn't be used for files and non-POSIX systems.
	Forced
	// Disabled force remove POSIX format.
	// Safe for all systems and writers.
	Disabled
)
