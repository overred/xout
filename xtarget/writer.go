package xtarget

import (
	"io"

	"github.com/overred/xout/xfield"
	"github.com/overred/xout/xformat"
	"github.com/overred/xout/xlevel"
)

type Writer struct {
	// output is an output destination.
	// Write operation will discard if output not defined.
	output io.Writer
	// level describes the Writer importance level.
	// It's helpful for Formatters.
	// Write operation will discard if level is zero.
	level xlevel.Level
	// fields contains some useful information.
	// It's helpful for Formatters.
	fields xfield.Fields
	// formatter is a preprocessor for data before it will
	// write to Output.
	// Data will write directly to Output if formatter not defined.
	formatter xformat.Formatter
}

// Write writes text into configured target.
func (w Writer) Write(p []byte) (int, error) {
	// Discard if Output not defined
	if w.output == nil {
		return 0, nil
	}
	// Discard if Level is zero
	if w.level == 0 {
		return 0, nil
	}

	// Use Formatter if defined
	if w.formatter != nil {
		// So io.MultiWriter has an implementation which returns an error
		// when length of given bytes not equal length bytes which were wrote
		// we use this hack.
		if n, err := w.formatter.Writer(w.output, w.level, w.fields).Write(p); err != nil {
			return n, err
		}
		return len(p), nil
	}
	return w.output.Write(p)
}
