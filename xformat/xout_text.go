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

// Writer creates new io.Writer to write into output through this formatter.
func (f Text) Writer(output io.Writer, level xlevel.Level, fields xfields.Fields) io.Writer {
	return TextWriter{
		output: output,
		level:  level,
		fields: fields,
	}
}

// TextWriter io.Writer implementation for this formatter.
type TextWriter struct {
	output io.Writer
	level  xlevel.Level
	fields xfields.Fields
}

// Write writes formatted data into output.
func (w TextWriter) Write(input []byte) (int, error) {
	if w.level == xlevel.Text {
		return w.output.Write(input)
	}

	levelName := strings.ToUpper(fmt.Sprintf("%-7s", w.level.Higher().String()))
	colorFormat := map[xlevel.Level]color.Color{
		xlevel.Trace: color.FgGray,
		xlevel.Debug: color.FgGray,
		xlevel.Info:  color.FgCyan,
		xlevel.Warn:  color.FgYellow,
		xlevel.Error: color.FgRed,
		xlevel.Fatal: color.FgRed,
		xlevel.Panic: color.FgRed,
	}[w.level.Higher()]
	if colorFormat == 0 {
		colorFormat = color.FgWhite
	}

	fields := strings.Builder{}
	for _, field := range w.fields.List() {
		fields.WriteString(colorFormat.Render(field.Name) + "=" + strconv.Quote(fmt.Sprint(field.Value)) + " ")
	}

	fieldsFormat := fields.String()
	if len(fieldsFormat) > 0 {
		fieldsFormat = fmt.Sprintf(colorFormat.Render("| %s"), fieldsFormat)
	}

	format := fmt.Sprintf(
		"%s %-45s %s\n",
		colorFormat.Render(fmt.Sprintf("%s %s |", time.Now().Format("15:04:05"), levelName)),
		string(strings.ReplaceAll(string(input), "\n", " ")),
		fieldsFormat,
	)
	return w.output.Write([]byte(format))
}
