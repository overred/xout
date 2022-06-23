package xout

import (
	"io"

	"github.com/overred/xout/xfields"
	"github.com/overred/xout/xlevel"
	"github.com/overred/xout/xtarget"
	"gopkg.in/gookit/color.v1"
)

// Logger describes logger and printer instance.
type Logger struct {
	// enable formatting tags support.
	tags bool
	// output targets.
	targets []xtarget.Target
	// fields contains some special data usually helpful for formatters.
	fields xfields.Fields
	// cache contains prepared writers.
	cache map[xlevel.Level]io.Writer
}

// writer describes io.Writer for XOut.
type writer struct {
	// output points to targets.
	output io.Writer
	// enable formatting tags support.
	tags bool
}

// New creates a new empty XOut logger.
// This is useful if you want to set up your own Targets.
// BE AWARE: This setup won't print anything,
// you need to bind targets.
func New() Logger {
	return Logger{}
}

// WithTags returns new instance with enabled/disabled formatting tags.
func (x Logger) WithTags(tags bool) Logger {
	x.tags = tags
	return x
}

// WithTarget returns new instance with new target added.
func (x Logger) WithTarget(target xtarget.Target) Logger {
	targets := make([]xtarget.Target, len(x.targets)+1)
	copy(targets, x.targets)
	targets[len(targets)-1] = target
	x.targets = targets
	x.cache = nil
	return x
}

// WithFields returns new instance with added fields.
func (x Logger) WithFields(fields xfields.Fields) Logger {
	x.fields = x.fields.Merge(fields)
	x.cache = nil
	return x
}

// WithField returns new instance with new field added.
func (x Logger) WithField(name string, value interface{}) Logger {
	x.fields = x.fields.With(name, value)
	x.cache = nil
	return x
}

// Writer returns writer to write into all targets.
func (x Logger) Writer(level xlevel.Level) io.Writer {
	var output io.Writer

	// Trying to find cache
	if x.cache == nil {
		x.cache = map[xlevel.Level]io.Writer{}
	} else {
		if wr, ok := x.cache[level]; ok {
			output = wr
		}
	}

	// Cache generation
	if output == nil {
		wrs := make([]io.Writer, 0, len(x.targets))
		for i := range x.targets {
			wrs = append(wrs, x.targets[i].Writer(level, x.fields))
		}
		output = io.MultiWriter(wrs...)
		x.cache[level] = output
	}

	return writer{
		output: output,
		tags:   x.tags,
	}
}

// Write writes data into all targets with preprocessing.
func (w writer) Write(p []byte) (int, error) {
	// Apply color tags rendering
	if w.tags {
		return w.output.Write([]byte(color.Render(string(p))))
	}
	return w.output.Write(p)
}
