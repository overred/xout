package xtarget

import (
	"io"
	"os"

	"github.com/mattn/go-colorable"
	"github.com/overred/xout/xcolor"
	"github.com/overred/xout/xfields"
	"github.com/overred/xout/xformat"
	"github.com/overred/xout/xlevel"
)

// Target describes target destination with some configurations.
type Target struct {
	// writer is an output destination. It may be console, file or any other stream.
	// If nil - output won't produced, and it isn't a error.
	writer io.Writer
	// level describes log level mask and acts like filter.
	// If zero - no one event will send to writer (silent mode).
	level xlevel.Level
	// color describes how must POSIX colors be processed for output writer.
	color xcolor.Mode
	// format describes specific formatter for this output.
	// If nil - raw data will be passed.
	// Formatter never gets Level(0) or LevelText so they are reserved.
	// Level(0) won't print and LevelText always will prints as a raw.
	format xformat.Formatter
	// fields describes additional information for debug methods.
	// Usually it's helpful for formatters.
	fields xfields.Fields
}

// New constructs output target.
func New() Target {
	return Target{}
}

// NewWriter preset creates pure target for given writer for all levels
// with automatic color mode and default formatter.
func NewWriter(writer io.Writer) Target {
	return New().
		WithWriter(writer).
		WithLevel(xlevel.Text.Elevate()).
		WithColor(xcolor.ModeAuto).
		WithFormat(xformat.LogrusText())
}

// NewStdout preset creates pure stdout writer for all levels
// with automatic color mode and default formatter.
func NewStdout() Target {
	return NewWriter(os.Stdout)
}

// NewStderr preset creates pure stderr writer for all levels
// with automatic color mode and default formatter.
func NewStderr() Target {
	return NewWriter(os.Stderr)
}

// WithWriter returns new Target with output target.
func (out Target) WithWriter(target io.Writer) Target {
	out.writer = target
	return out
}

// Writer returns current writer.
func (out Target) Writer() io.Writer {
	return out.writer
}

// WithLevel returns new Target with log levels.
func (out Target) WithLevel(level xlevel.Level) Target {
	out.level = level
	return out
}

// AddLevel returns new Target with added log level to existing.
func (out Target) AddLevel(level xlevel.Level) Target {
	out.level |= level
	return out
}

// Level returns current level mask.
func (out Target) Level() xlevel.Level {
	return out.level
}

// WithColor returns new Target with color mode.
func (out Target) WithColor(color xcolor.Mode) Target {
	out.color = color
	return out
}

// Color returns current color mode.
func (out Target) Color() xcolor.Mode {
	return out.color
}

// WithFormat returns new Target with formatter.
func (out Target) WithFormat(format xformat.Formatter) Target {
	out.format = format
	return out
}

// Format returns current formatter.
func (out Target) Format() xformat.Formatter {
	return out.format
}

// WithFields returns new Target with fields.
func (out Target) WithFields(fields xfields.Fields) Target {
	out.fields = fields
	return out
}

// AddField returns new Target with updated fields.
func (out Target) AddField(field xfields.Field) Target {
	out.fields = out.fields.AddField(field)
	return out
}

// AddNameVal returns new Target with updated fields.
func (out Target) AddNameVal(name string, value interface{}) Target {
	out.fields = out.fields.AddNameVal(name, value)
	return out
}

// Fields returns current fields.
func (out Target) Fields() xfields.Fields {
	return out.fields
}

// Write sends data into configured writer.
// Before send it checks level, processes colors and calls formatter.
// Fields presents additional information generally for formatters and can be nil.
func (out Target) Write(level xlevel.Level, input string) error {
	// Skip if no destination
	if out.writer == nil {
		return nil
	}
	// Skip if zero level (silent mode)
	if out.level.Has(0) {
		return nil
	}
	// Skip if configured level mask not contain given level
	if !out.level.Has(level) {
		return nil
	}

	// If writer can be cast to *os.File - wrap it by Windows safe writer
	var writer io.Writer
	if f, ok := out.writer.(*os.File); ok {
		writer = colorable.NewColorable(f)
	} else {
		writer = out.writer
	}

	// Decision about colorful output
	switch out.color {
	case xcolor.ModeDisabled:
		// Disable colors at all
		writer = colorable.NewNonColorable(writer)
	case xcolor.ModeAuto:
		// If destination not a file - disable colors
		if _, ok := out.writer.(*os.File); !ok {
			writer = colorable.NewNonColorable(writer)
		} else
		// If defined special non-color variables - disable colors
		if _, no_color := os.LookupEnv("NO_COLOR"); no_color || os.Getenv("CLICOLOR") == "0" {
			writer = colorable.NewNonColorable(writer)
		} else
		// In other cases - leave colors as is
		{
			writer = out.writer
		}
	default:
		// In case of any unknown color mode - pass it as is
		writer = out.writer
	}

	// If formatter exists and it is log, not a regular text - pass through
	if level != xlevel.Text && out.format != nil {
		_, err := writer.Write([]byte(out.format(level, input, out.fields)))
		return err
	}

	// Pass input as is if no formatter
	_, err := writer.Write([]byte(input))
	return err
}
