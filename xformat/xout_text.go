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

// Text basic formatter with log colorization and fields support.
type Text struct{}

// NewText creates new basic formatter.
func NewText() Text {
	return Text{}
}

// TextWriter io.Writer implementation for this formatter.
type TextWriter struct {
	output io.Writer
	level  xlevel.Level
	fields xfields.Fields
	color  color.Color
	format string
}

// Writer creates new io.Writer to write into output through this formatter.
func (f Text) Writer(output io.Writer, level xlevel.Level, fields xfields.Fields) io.Writer {
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

	formatLevel := c.Render(strings.ToUpper(fmt.Sprintf("%-7s", level.Higher().String())))

	formatFields := strings.Builder{}
	formatFields.Grow(1 << 12)
	for i := 0; i < fields.Count(); i++ {
		field := fields.Index(i)
		formatFields.WriteString(c.Render(field.Name) + "=" + strconv.Quote(fmt.Sprint(field.Value)) + " ")
	}

	format := fmt.Sprintf("%%s %s %s %%-45s",
		formatLevel,
		c.Render("|"),
	)
	if formatFields.Len() == 0 {
		format += "\n"
	} else {
		format += c.Render(" | ") + formatFields.String() + "\n"
	}

	return TextWriter{
		output: output,
		level:  level,
		fields: fields,
		color:  c,
		format: format,
	}
}

// Write writes formatted data into output.
func (w TextWriter) Write(input []byte) (int, error) {
	if w.level == xlevel.Text {
		return w.output.Write(input)
	}
	format := fmt.Sprintf(
		w.format,
		w.color.Render(time.Now().Format("15:04:05")),
		string(input),
	)
	return w.output.Write([]byte(format))
}
