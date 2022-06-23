package xformat

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/overred/xout/xfields"
	"github.com/overred/xout/xlevel"
	"gopkg.in/gookit/color.v1"
)

// LogrusText logrus-like text formatter.
type LogrusText struct {
	Start time.Time
}

// NewLogrusText creates new formatter.
func NewLogrusText() LogrusText {
	return LogrusText{
		Start: time.Now(),
	}
}

// LogrusTextWriter io.Writer implementation for this formatter.
type LogrusTextWriter struct {
	output io.Writer
	level  xlevel.Level
	start  time.Time
	format string
}

// Writer creates new io.Writer to write into output through this formatter.
func (f LogrusText) Writer(output io.Writer, level xlevel.Level, fields xfields.Fields) io.Writer {
	c := map[xlevel.Level]color.Color{
		xlevel.Trace: color.FgGray,
		xlevel.Debug: color.FgGray,
		xlevel.Info:  color.FgCyan,
		xlevel.Warn:  color.FgYellow,
		xlevel.Error: color.FgRed,
		xlevel.Fatal: color.FgRed,
		xlevel.Panic: color.FgRed,
	}[level]
	if c == 0 {
		c = color.FgWhite
	}

	formatFields := strings.Builder{}
	formatFields.Grow(1 << 12)
	for i := 0; i < fields.Count(); i++ {
		field := fields.Index(i)
		formatFields.WriteString(c.Render(field.Name) + "=" + strconv.Quote(field.String()) + " ")
	}

	format := fmt.Sprintf("%s[%%04d] %%-45s %s\n",
		c.Render(strings.ToUpper(fmt.Sprintf("%-4s", level.Higher().String()))),
		formatFields.String(),
	)

	return LogrusTextWriter{
		output: output,
		level:  level,
		start:  f.Start,
		format: format,
	}
}

// Write writes formatted data into output.
func (w LogrusTextWriter) Write(input []byte) (int, error) {
	if w.level == xlevel.Text {
		return w.output.Write(input)
	}
	format := fmt.Sprintf(
		w.format,
		int(time.Since(w.start).Seconds()),
		strings.ReplaceAll(string(input), "\n", " "),
	)
	return w.output.Write([]byte(format))
}
