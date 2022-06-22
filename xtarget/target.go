package xtarget

import (
	"io"
	"os"

	"github.com/mattn/go-colorable"
	"github.com/overred/xout/xfields"
	"github.com/overred/xout/xformat"
	"github.com/overred/xout/xlevel"
	"github.com/overred/xout/xposix"
)

// Target describes target destination with some configurations.
type Target struct {
	// Output is an output destination. It may be console, file or any other stream.
	// If nil - output won't produced.
	Output io.Writer
	// LevelMask describes levels mask and acts like filter.
	// If zero - no one event will send to writer (silent mode).
	LevelMask xlevel.Level
	// PosixMode describes how must POSIX symbols be processed
	// before send to Output. Zero means automatic mode.
	PosixMode xposix.Mode
	// Formatter describes specific formatter for this output.
	// If nil - raw data will be passed.
	Formatter xformat.Formatter
}

// New creates new instance of Target.
// LevelMask for all levels, automatic PosixMode, and
// Default formatter.
func New(output io.Writer) Target {
	return Target{
		Output:    output,
		LevelMask: xlevel.Text.Elevate(),
		PosixMode: xposix.Auto,
		Formatter: xformat.NewDefault(),
	}
}

// WithOutput returns a copy of Target with new Output.
func (t Target) WithOutput(output io.Writer) Target {
	t.Output = output
	return t
}

// WithLevelMask returns a copy of Target with new LevelMask.
func (t Target) WithLevelMask(mask xlevel.Level) Target {
	t.LevelMask = mask
	return t
}

// WithPosixMode returns a copy of Target with new PosixMode.
func (t Target) WithPosixMode(mode xposix.Mode) Target {
	t.PosixMode = mode
	return t
}

// WithFormatter returns a copy of Target with new Formatter.
func (t Target) WithFormatter(formatter xformat.Formatter) Target {
	t.Formatter = formatter
	return t
}

// Writer produces new io.Writer compatible object according Target's configuration.
// Fields are optional and usually helpful for formatters.
func (t Target) Writer(level xlevel.Level, fields xfields.Fields) io.Writer {
	// Skip if no destination
	if t.Output == nil {
		return io.Discard
	}
	// Skip if zero level (silent mode)
	if t.LevelMask.Has(0) {
		return io.Discard
	}
	// Skip if configured level mask not contain given level
	if !t.LevelMask.Has(level) {
		return io.Discard
	}

	// If output can be cast to *os.File - wrap it by Windows safe output
	var output io.Writer
	if f, ok := t.Output.(*os.File); ok {
		output = colorable.NewColorable(f)
	} else {
		output = t.Output
	}

	// Decision about colorful output
	switch t.PosixMode {
	case xposix.Disabled:
		// Disable colors at all
		output = colorable.NewNonColorable(output)
	case xposix.Auto:
		// If destination not a file - disable colors
		if _, ok := t.Output.(*os.File); !ok {
			output = colorable.NewNonColorable(output)
		} else
		// If defined special non-color variables - disable colors
		if _, no_color := os.LookupEnv("NO_COLOR"); no_color || os.Getenv("CLICOLOR") == "0" {
			output = colorable.NewNonColorable(output)
		} else
		// In other cases - leave colors as is
		{
			output = t.Output
		}
	default:
		// In case of any unknown color mode - pass it as is
		output = t.Output
	}

	return Writer{
		output:    output,
		level:     level,
		fields:    fields,
		formatter: t.Formatter,
	}
}