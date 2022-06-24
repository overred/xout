package xout

import (
	"io"

	"github.com/overred/xout/xcaller"
	"github.com/overred/xout/xfields"
	"github.com/overred/xout/xlevel"
	"github.com/overred/xout/xtarget"
	"gopkg.in/gookit/color.v1"
)

// Logger describes logger and printer instance.
type Logger struct {
	// enable formatting tags support.
	tags bool
	// caller add caller information into fields.
	caller bool
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

// WithCaller returns new instance with enabled/disabled formatting tags.
// BE AWARE: Enabled caller info will flushes cache every time when log calls.
// So this option has a bad value for logger performance. Use only for debugging.
func (x Logger) WithCaller(caller bool) Logger {
	x.caller = caller
	return x
}

// WithTarget returns new instance with new target added.
func (x Logger) WithTarget(target ...xtarget.Target) Logger {
	targets := make([]xtarget.Target, len(x.targets)+len(target))
	copy(targets, x.targets)
	for i := 0; i < len(target); i++ {
		targets[len(targets)-1-i] = target[i]
	}
	x.targets = targets
	x.cache = nil
	return x
}

// WithFields returns new instance with added fields set or multiple sets.
func (x Logger) WithFields(fields ...xfields.Fields) Logger {
	for i := range fields {
		x.fields = x.fields.Merge(fields[i])
	}
	x.cache = nil
	return x
}

// WithField returns new instance with new field added.
func (x Logger) WithField(name string, value interface{}) Logger {
	x.fields = x.fields.With(name, value)
	x.cache = nil
	return x
}

// WithError returns new instance with new error field added.
func (x Logger) WithError(err error) Logger {
	x.fields = x.fields.With("error", err.Error())
	x.cache = nil
	return x
}

// Writer returns writer to write into all targets.
func (x Logger) Writer(level xlevel.Level) io.Writer {
	var output io.Writer

	// If caller's info enabled attach additional fields.
	// It isn't possible to use cache, so fields changes frequently.
	if x.caller {
		// frame depth it's number of function calls from any Logger.LogFn() to xcaller.Lookup().
		// Think about it as a same magic number.
		const frameDepth = 4
		x = x.WithFields(xcaller.AsFields(frameDepth))
	} else {
		// Trying to use cache
		if x.cache == nil {
			x.cache = map[xlevel.Level]io.Writer{}
		} else {
			if wr, ok := x.cache[level]; ok {
				output = wr
			}
		}
	}

	// Creating multi-writer for targets if didn't find
	// in cache or caller info enabled.
	// If caller info enabled it isn't possible to use cache,
	// so fields changes every time.
	if output == nil {
		wrs := make([]io.Writer, 0, len(x.targets))
		for i := range x.targets {
			wrs = append(wrs, x.targets[i].Writer(level, x.fields))
		}
		output = io.MultiWriter(wrs...)
		// No reasons for caching.
		if !x.caller {
			x.cache[level] = output
		}
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
