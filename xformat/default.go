package xformat

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/overred/xout/xfields"
	"github.com/overred/xout/xlevel"
	"gopkg.in/gookit/color.v1"
)

// Default basic formatter with log colorization and fields support.
type Default struct {
	startTime time.Time
}

// NewDefault creates new basic formatter.
func NewDefault() Default {
	return Default{
		startTime: time.Now(),
	}
}

// Writer creates new io.Writer to write into output through this formatter.
func (f Default) Writer(output io.Writer, level xlevel.Level, fields xfields.Fields) io.Writer {
	return DefaultWriter{
		output:    output,
		level:     level,
		fields:    fields,
		startTime: f.startTime,
	}
}

// DefaultWriter io.Writer implementation for this formatter.
type DefaultWriter struct {
	output    io.Writer
	level     xlevel.Level
	fields    xfields.Fields
	startTime time.Time
}

// Write writes formatted data into output.
func (w DefaultWriter) Write(input []byte) (int, error) {
	// Pass text level as is
	if w.level == xlevel.Text {
		return w.output.Write(input)
	}

	levelName := strings.ToUpper(fmt.Sprintf("%-7s", w.level.Higher().String()))
	colorFormat := map[xlevel.Level]color.Color{
		xlevel.Trace: color.FgGray,
		xlevel.Debug: color.FgGray,
		xlevel.Info:  color.FgBlue,
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
		fields.WriteString(colorFormat.Render(field.Name) + "=" + field.String() + " ")
	}

	format := fmt.Sprintf(
		"%s %s %-44s %s\n",
		colorFormat.Render(fmt.Sprintf("%04d", int(time.Since(w.startTime).Seconds()))),
		colorFormat.Render(levelName),
		string(strings.ReplaceAll(string(input), "\n", " ")),
		fields.String(),
	)
	return w.output.Write([]byte(format))
}
