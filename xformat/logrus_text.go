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

// LogrusText logrus-text like formatter.
type LogrusText struct {
	startTime time.Time
}

// NewLogrusText creates new logrus-text like formatter.
func NewLogrusText() LogrusText {
	return LogrusText{
		startTime: time.Now(),
	}
}

// Writer creates new io.Writer to write into output through this formatter.
func (f LogrusText) Writer(output io.Writer, level xlevel.Level, fields xfields.Fields) io.Writer {
	return LogrusTextWriter{
		output:    output,
		level:     level,
		fields:    fields,
		startTime: f.startTime,
	}
}

// LogrusTextWriter io.Writer implementation for this formatter.
type LogrusTextWriter struct {
	output    io.Writer
	level     xlevel.Level
	fields    xfields.Fields
	startTime time.Time
}

// Write writes formatted data into output.
func (w LogrusTextWriter) Write(input []byte) (int, error) {
	// Pass text level as is
	if w.level == xlevel.Text {
		return w.output.Write(input)
	}

	levelName := strings.ToUpper(fmt.Sprintf("%.4s", w.level.Higher().String()))
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

	// Original logrus text formatter has reverted values sequence
	fields := strings.Builder{}
	list := w.fields.List()
	for i := len(list) - 1; i >= 0; i-- {
		fields.WriteString(colorFormat.Render(list[i].Name) + "=" + list[i].String() + " ")
	}

	format := fmt.Sprintf(
		"%-4s[%04d] %-45s %s\n",
		colorFormat.Render(levelName),
		int(time.Since(w.startTime).Seconds()),
		string(strings.ReplaceAll(string(input), "\n", " ")),
		fields.String(),
	)
	return w.output.Write([]byte(format))
}
