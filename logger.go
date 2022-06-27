package xout

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"

	"github.com/overred/xout/xfield"
	"github.com/overred/xout/xlevel"
	"github.com/overred/xout/xtarget"
	"gopkg.in/gookit/color.v1"
)

var (
	// Default contains the configuration of the standard logger.
	// It will be used for functions called without creating a logger instance.
	// This value can be overridden and will have an impact on the entire project.
	// Redefinition can be useful in the case of microservices.
	Default = NewPresetDefault()
)

// Logger describes the logger object.
// This is a more flexible alternative to the built-in Default logger.
type Logger struct {
	// tags flag for parsing formatting tags.
	// If enabled, the tags supplied by the gookit/color module will be processed.
	// Has a strong impact on performance.
	tags bool
	// caller flag for adding a call function.
	// If enabled, the file, line, and name of the caller will be added to the log information.
	// It has a strong impact on performance, because the cache will be disabled.
	// Debugging purposes only recommended.
	caller bool
	// targets describes a list of targets for outputting logs and texts.
	// It is possible to define the format and logging levels for each.
	targets []xtarget.Target
	// fields contain additional information.
	// May include a description of errors, caller info, or any user information.
	fields xfield.Fields
	// cache contains prepared writers for reusing.
	// The cache won't work if caller flag enabled.
	// The cache flushes every time the fields or targets change.
	cache *sync.Map
	// exit contains function to be used instead of os.Exit() which callings with Fatal & Fatalf methods.
	exit func(code int)
}

// writer provides io.Writer interface.
// Also it parses color tags if enabled.
type writer struct {
	// output points to targets.
	output io.Writer
	// enable formatting tags support.
	tags bool
}

// New creates new Logger instance.
// This is a more flexible alternative to the built-in Default logger.
// Logger produced by the New() won't print anything, you should configure targets first
// or use presets.
func New() Logger {
	return Logger{
		cache: &sync.Map{},
		exit:  os.Exit,
	}
}

// WithTags returns a new logger instance with the gookit/color formatting tag parsing flag.
// The cache will be reset for the new logger instance.
// Parsing formatting tags has a strong impact on performance.
// It is recommended to disable this option for cases where high performance is required.
func (x Logger) WithTags(tags bool) Logger {
	x.tags = tags
	x.cache = &sync.Map{}
	return x
}

// WithTags returns a new logger instance with the gookit/color formatting tag parsing flag.
// The cache will be reset for the new logger instance.
// Parsing formatting tags has a strong impact on performance.
// It is recommended to disable this option for cases where high performance is required.
func WithTags(tags bool) Logger {
	return Default.WithTags(tags)
}

// WithCaller returns a new instance of the logger with the flag of the calling function info.
// This option adds a new field to the information fields for each log call.
// The logger cache will be disabled, which will have a strong impact on performance.
func (x Logger) WithCaller(caller bool) Logger {
	x.caller = caller
	return x
}

// WithCaller returns a new instance of the logger with the flag of the calling function info.
// This option adds a new field to the information fields for each log call.
// The logger cache will be disabled, which will have a strong impact on performance.
func WithCaller(caller bool) Logger {
	return Default.WithCaller(caller)
}

// WithTargets returns a new logger instance and adds a new target to it.
// The cache will be reset for the new logger instance.
// It is recommended to set targets once before using the logger.
func (x Logger) WithTargets(target ...xtarget.Target) Logger {
	targets := make([]xtarget.Target, len(x.targets)+len(target))
	copy(targets, x.targets)
	for i := 0; i < len(target); i++ {
		targets[len(targets)-1-i] = target[i]
	}
	x.targets = targets
	x.cache = &sync.Map{}
	return x
}

// WithTargets returns a new logger instance and adds a new target to it.
// The cache will be reset for the new logger instance.
// It is recommended to set targets once before using the logger.
func WithTargets(target ...xtarget.Target) Logger {
	return Default.WithTargets(target...)
}

// WithFields returns a new instance of the logger and adds fields to it.
// The cache of the new logger will be reset.
// For performance-critical situations, it is recommended to set the fields before using the logger.
func (x Logger) WithFields(fields ...xfield.Fields) Logger {
	for i := range fields {
		x.fields = x.fields.Merge(fields[i])
	}
	x.cache = &sync.Map{}
	return x
}

// WithFields returns a new instance of the logger and adds fields to it.
// The cache of the new logger will be reset.
// For performance-critical situations, it is recommended to set the fields before using the logger.
func WithFields(fields ...xfield.Fields) Logger {
	return Default.WithFields(fields...)
}

// WithField returns a new instance of the logger and adds a field to it.
// The cache of the new logger will be reset.
// For performance-critical situations, it is recommended to set the fields before using the logger.
func (x Logger) WithField(name string, value interface{}) Logger {
	x.fields = x.fields.With(name, value)
	x.cache = &sync.Map{}
	return x
}

// WithField returns a new instance of the logger and adds a field to it.
// The cache of the new logger will be reset.
// For performance-critical situations, it is recommended to set the fields before using the logger.
func WithField(name string, value interface{}) Logger {
	return Default.WithField(name, value)
}

// WithError returns a new instance of the logger and adds a field with an error to it.
// If nil passed then this field will be removed.
// The cache of the new logger will be reset.
// For performance-critical situations, it is recommended to set the fields before using the logger.
func (x Logger) WithError(err error) Logger {
	if err == nil {
		x.fields = x.fields.Remove(xfield.NameError)
	} else {
		x.fields = x.fields.With(xfield.NameError, err.Error())
	}
	x.cache = &sync.Map{}
	return x
}

// WithError returns a new instance of the logger and adds a field with an error to it.
// If nil passed then this field will be removed.
// The cache of the new logger will be reset.
// For performance-critical situations, it is recommended to set the fields before using the logger.
func WithError(err error) Logger {
	return Default.WithError(err)
}

// WithExit returns a new instance of the logger and sets the fatal error handling function.
// This can be useful in cases where it is necessary to handle an error before terminating the program.
func (x Logger) WithExit(exit func(code int)) Logger {
	x.exit = exit
	return x
}

// WithExit returns a new instance of the logger and sets the fatal error handling function.
// This can be useful in cases where it is necessary to handle an error before terminating the program.
func WithExit(exit func(code int)) Logger {
	return Default.WithExit(exit)
}

// Writer prepares the data processing stream and returns an instance of io.Writer for the specified logging level.
// The returned io.Writer is cached if possible.
// The data sent to the returned io.Writer will be transmitted and processed by all targets.
func (x Logger) Writer(level xlevel.Level) io.Writer {
	var output io.Writer

	// If caller's info enabled attach additional fields.
	// It isn't possible to use cache, so fields changes frequently.
	if x.caller {
		// frameDepth it's number of function calls from any Logger.LogFn() to xcaller.Lookup().
		// Think about it as a same magic number.
		const frameDepth = 2
		if pc, file, line, ok := runtime.Caller(frameDepth); ok {
			fn := runtime.FuncForPC(pc)
			x = x.WithField(
				xfield.NameCaller,
				fmt.Sprintf("%s:%d %s()", file, line, fn.Name()),
			)
		}
	} else {
		// Try to use cache
		if val, ok := x.cache.Load(level); ok {
			output = val.(io.Writer)
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
		// No reasons for caching if report caller enabled.
		if !x.caller {
			x.cache.Store(level, output)
		}
	}

	// Decision with zero allocation for simple data without processing.
	if x.tags {
		return writer{
			output: output,
			tags:   x.tags,
		}
	}
	return output
}

// Writer prepares the data processing stream and returns an instance of io.Writer for the specified logging level.
// The returned io.Writer is cached if possible.
// The data sent to the returned io.Writer will be transmitted and processed by all targets.
func Writer(level xlevel.Level) io.Writer {
	return Default.Writer(level)
}

// Write sends data to all targets.
// Parses formatting tags if they are enabled.
// Fields and logging level are prepared earlier using Logger.Writer()
func (w writer) Write(p []byte) (int, error) {
	// Apply color tags rendering
	if w.tags {
		return w.output.Write([]byte(color.Render(string(p))))
	}
	return w.output.Write(p)
}
