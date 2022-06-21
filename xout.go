package xout

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/overred/xout/xfields"
	"github.com/overred/xout/xlevel"
	"github.com/overred/xout/xtarget"
	"gopkg.in/gookit/color.v1"
)

// Xout describes logger and printer instance.
type Xout struct {
	// enable formatting tags support.
	tags bool
	// output targets.
	targets []xtarget.Target
	// fields contains some special data usually helpful for formatters.
	fields xfields.Fields
}

// New creates new instance of Xout.
func New() Xout {
	return Xout{}
}

// NewDefault creates new instance with default targets.
// Levels xlevel.Error, xlevel.Fatal, xlevel.Panic goes to os.Stderr.
// All other levels goes to os.Stdout.
// Automatic color mode for all targets with formatting tags enabled.
func NewDefault() Xout {
	return New().
		WithTags(true).
		AddTarget(
			xtarget.NewStdout().WithLevel(
				xlevel.Text | xlevel.Trace | xlevel.Debug | xlevel.Info | xlevel.Warn,
			),
		).
		AddTarget(
			xtarget.NewStderr().WithLevel(
				xlevel.Error | xlevel.Fatal | xlevel.Panic,
			),
		)
}

// WithTags returns new instance with enabled/disabled formatting tags.
func (x Xout) WithTags(tags bool) Xout {
	x.tags = tags
	return x
}

// Tags returns current formatting tags status.
func (x Xout) Tags() bool {
	return x.tags
}

// WithTargets returns new instance with replaced output targets.
func (x Xout) WithTargets(targets []xtarget.Target) Xout {
	x.targets = targets
	return x
}

// AddTarget returns new instance with added target.
func (x Xout) AddTarget(target xtarget.Target) Xout {
	targets := make([]xtarget.Target, len(x.targets)+1)
	copy(targets, x.targets)
	targets[len(targets)-1] = target
	x.targets = targets
	return x
}

// Targets returns current targets.
func (x Xout) Targets() []xtarget.Target {
	return x.targets
}

// WithFields returns new instance with replaced fields.
func (x Xout) WithFields(fields xfields.Fields) Xout {
	x.fields = fields
	return x
}

// AddField returns new instance with new field.
func (x Xout) AddField(field xfields.Field) Xout {
	x.fields = x.fields.AddField(field)
	return x
}

// AddNameVal returns new instance with new field.
func (x Xout) AddNameVal(name string, value interface{}) Xout {
	x.fields = x.fields.
		AddField(
			xfields.
				NewField().
				WithName(name).
				WithValue(value),
		)
	return x
}

// Write writes to all configured targets which are fit by level.
func (x Xout) Write(level xlevel.Level, input string) error {
	// If formatting tags enabled - parse and convert to POSIX
	if x.tags {
		buff := bytes.NewBuffer([]byte{})
		color.Fprintf(buff, input)
		data, _ := ioutil.ReadAll(buff)
		input = string(data)
	}

	errs := make([]error, 0, len(x.targets))
	for i := range x.targets {
		err := x.targets[i].WithFields(x.fields).Write(level, input)
		if err != nil {
			errs = append(errs, err)
		}
	}

	// Take very first error if at least one happened
	if len(errs) != 0 {
		return errs[0]
	}
	return nil
}

// Printf formats according to a format specifier and returns the resulting string.
func (x Xout) Printf(format string, v ...interface{}) error {
	return x.Write(xlevel.Text, fmt.Sprintf(format, v...))
}

// Print formats using the default formats for its operands and writes to standard output.
// Unlike fmt.Print() don't add spaces between elements.
func (x Xout) Print(a ...interface{}) error {
	return x.Printf(strings.Repeat("%v", len(a)), a...)
}

// Println formats using the default formats for its operands and writes
// to standard output with new line at the end.
// Unlike fmt.Print() don't add spaces between elements.
func (x Xout) Println(a ...interface{}) error {
	args := make([]interface{}, len(a), len(a)+1)
	copy(args, a)
	args = append(args, "\n")
	return x.Print(args...)
}

// Tracef formats and writes trace log to targets.
func (x Xout) Tracef(format string, v ...interface{}) error {
	return x.Write(xlevel.Trace, fmt.Sprintf(format, v...))
}

// Debugf formats and writes debug log to targets.
func (x Xout) Debugf(format string, v ...interface{}) error {
	return x.Write(xlevel.Debug, fmt.Sprintf(format, v...))
}

// Infof formats and writes info log to targets.
func (x Xout) Infof(format string, v ...interface{}) error {
	return x.Write(xlevel.Info, fmt.Sprintf(format, v...))
}

// Warnf formats and writes warn log to targets.
func (x Xout) Warnf(format string, v ...interface{}) error {
	return x.Write(xlevel.Warn, fmt.Sprintf(format, v...))
}

// Errorf formats and writes error log to targets.
func (x Xout) Errorf(format string, v ...interface{}) error {
	return x.Write(xlevel.Error, fmt.Sprintf(format, v...))
}

// Fatalf formats and writes fatal log to targets.
// BE AWARE: This operation calls os.Exit().
func (x Xout) Fatalf(format string, v ...interface{}) error {
	x.Write(xlevel.Fatal, fmt.Sprintf(format, v...))
	os.Exit(1)
	return nil
}

// Panicf formats and writes panic log to targets.
// BE AWARE: This operation calls panic().
func (x Xout) Panicf(format string, v ...interface{}) error {
	x.Write(xlevel.Panic, fmt.Sprintf(format, v...))
	panic(fmt.Sprintf(format, v...))
}
