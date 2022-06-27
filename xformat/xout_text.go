package xformat

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/overred/xout/xfield"
	"github.com/overred/xout/xlevel"
	"gopkg.in/gookit/color.v1"
)

// textFormatter basic formatter with log colorization and fields support.
type textFormatter struct {
	color map[xlevel.Level]color.Color
}

// NewText creates new basic formatter.
func NewText() Formatter {
	return textFormatter{
		color: map[xlevel.Level]color.Color{
			xlevel.Trace: color.FgGray,
			xlevel.Debug: color.FgGray,
			xlevel.Info:  color.FgCyan,
			xlevel.Warn:  color.FgYellow,
			xlevel.Error: color.FgRed,
			xlevel.Fatal: color.FgRed,
			xlevel.Panic: color.FgRed,
		},
	}
}

// textWriter io.Writer implementation for this formatter.
type textWriter struct {
	output io.Writer
	level  xlevel.Level
	fields xfield.Fields
	color  color.Color
	format string
}

// Writer creates new io.Writer to write into output through this formatter.
func (f textFormatter) Writer(output io.Writer, level xlevel.Level, fields xfield.Fields) io.Writer {
	if level == xlevel.Text {
		return output
	}

	c := f.color[level]
	if c == 0 {
		c = color.FgWhite
	}

	// ! Low-Performance Operation: 10 allocations
	formatLevel := c.Render(strings.ToUpper(fmt.Sprintf("%-7s", level.Higher().String())))

	formatFields := strings.Builder{}
	formatFields.Grow(1 << 8)
	for i := 0; i < fields.Count(); i++ {
		field := fields.Index(i)
		// ! Low-Performance Operation: About 8 allocations per field
		formatFields.WriteString(c.Render(field.Name) + "=" + strconv.Quote(fmt.Sprint(field.Value)) + " ")
	}

	// ! Low-Performance Operation: 8 allocations
	format := fmt.Sprintf("%%s %s | %%-45s",
		formatLevel,
	)
	if formatFields.Len() == 0 {
		format += "\n"
	} else {
		format += " | " + formatFields.String() + "\n"
	}

	return textWriter{
		output: output,
		level:  level,
		fields: fields,
		color:  c,
		format: format,
	}
}

// Write writes formatted data into output.
func (w textWriter) Write(input []byte) (int, error) {
	// ! Low-Performance Operation: 7 allocations
	format := fmt.Sprintf(
		w.format,
		time.Now().Format("15:04:05"),
		string(input),
	)
	return w.output.Write([]byte(format))
}
