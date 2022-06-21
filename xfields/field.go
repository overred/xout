package xfields

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/mattn/go-colorable"
)

// Field contains some specific field info for formatters.
type Field struct {
	// name contains field name.
	name string
	// value contains field value.
	value interface{}
}

// NewField creates new field container.
func NewField() Field {
	return Field{}
}

// WithName returns new field with specific name.
func (f Field) WithName(name string) Field {
	f.name = name
	return f
}

// Name returns current field's name.
func (f Field) Name() string {
	return f.name
}

// WithValue returns new field with specific value.
func (f Field) WithValue(value interface{}) Field {
	f.value = value
	return f
}

// Value returns current field's value.
func (f Field) Value() interface{} {
	return f.value
}

// String returns current field's value wrapped into
// quotes and escaped if necessary.
func (f Field) String() string {
	value := fmt.Sprintf("%v", f.value)

	// Remove POSIX formatting
	buff := bytes.NewBuffer([]byte{})
	colorable.NewNonColorable(buff).Write([]byte(value))
	plain, _ := ioutil.ReadAll(buff)
	value = string(plain)

	if strings.Contains(value, " ") {
		value = strconv.Quote(value)
	}
	return value
}
