package xformat

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/overred/xout/xfields"
	"github.com/overred/xout/xlevel"
)

// LogrusJson logrus-json like formatter.
type LogrusJson struct{}

// NewLogrusJson creates new logrus-json like formatter.
func NewLogrusJson() LogrusJson {
	return LogrusJson{}
}

// Writer creates new io.Writer to write into output through this formatter.
func (f LogrusJson) Writer(output io.Writer, level xlevel.Level, fields xfields.Fields) io.Writer {
	return LogrusJsonWriter{
		output: output,
		level:  level,
		fields: fields,
	}
}

// LogrusJsonWriter io.Writer implementation for this formatter.
type LogrusJsonWriter struct {
	output io.Writer
	level  xlevel.Level
	fields xfields.Fields
}

// Write writes formatted data into output.
func (w LogrusJsonWriter) Write(input []byte) (int, error) {
	if w.level == xlevel.Text {
		return w.output.Write(input)
	}

	object := map[string]interface{}{}
	object["level"] = w.level.Higher().String()
	object["msg"] = string(input)
	object["time"] = time.Now().Format(time.RFC3339)

	for _, field := range w.fields.List() {
		if _, exists := object[field.Name]; exists {
			object["fields."+field.Name] = field.Value
		} else {
			object[field.Name] = field.Value
		}
	}

	data, _ := json.Marshal(object)
	return w.output.Write([]byte(fmt.Sprintf("%s\n", string(data))))
}
