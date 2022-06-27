package xformat

import (
	"io"

	"github.com/overred/xout/xfield"
	"github.com/overred/xout/xlevel"
)

// fastTextFormatter basic formatter with log colorization and fields support.
type fastTextFormatter struct{}

// NewFastText creates new basic formatter.
func NewFastText() Formatter {
	return fastTextFormatter{}
}

// TextWriter io.Writer implementation for this formatter.
type fastTextWriter struct {
	output     io.Writer
	level      xlevel.Level
	formatPre  []byte
	formatPost []byte
}

// Writer creates new io.Writer to write into output through this formatter.
func (f fastTextFormatter) Writer(output io.Writer, level xlevel.Level, fields xfield.Fields) io.Writer {
	if level == xlevel.Text {
		return output
	}

	formatLevel := []byte(level.Higher().String())

	formatFields := make([]byte, 0, 1<<8)
	for i := 0; i < fields.Count(); i++ {
		field := fields.Index(i)
		formatFields = append(formatFields, []byte(field.Name)...)
		formatFields = append(formatFields, '=')
		formatFields = append(formatFields, []byte(field.String())...)
		formatFields = append(formatFields, ' ')
	}

	formatPre := make([]byte, 0, 1<<8)
	formatPre = append(formatPre, formatLevel...)
	formatPre = append(formatPre, ':', ' ')

	formatPost := make([]byte, 0, 1<<8)
	formatPost = append(formatPost, ' ')
	formatPost = append(formatPost, formatFields...)
	formatPost = append(formatPost, '\n')

	return fastTextWriter{
		output:     output,
		level:      level,
		formatPre:  formatPre,
		formatPost: formatPost,
	}
}

// Write writes formatted data into output.
func (w fastTextWriter) Write(input []byte) (int, error) {
	format := make([]byte, 0, len(w.formatPre)+len(input)+len(w.formatPost))
	format = append(format, w.formatPre...)
	format = append(format, input...)
	format = append(format, w.formatPost...)
	return w.output.Write(format)
}
