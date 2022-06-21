package xcolor

// Mode describes how must POSIX colors be processed for output writer.
type Mode uint8

const (
	// ModeAuto tries to use colors where it possible.
	// This option will remove colors if NO_COLOR env var exists
	// or CLICOLOR=0 or if output target isn't *io.File.
	// This is Windows safe and colored option.
	ModeAuto Mode = iota
	// ModeForced pass colors as is.
	// This option will look fine for Windows also
	// but it shouldn't be used for files.
	ModeForced
	// ModeDisabled force remove colors.
	// Safe for all systems and writers.
	ModeDisabled
)
