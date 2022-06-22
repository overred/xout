package xfields

import (
	"fmt"
	"strconv"
	"strings"
)

// Field single field with some information for formatters.
type Field struct {
	Name  string
	Value interface{}
}

// String converts value to string and quotes it if necessary.
func (f Field) String() string {
	str := fmt.Sprintf("%v", f.Value)
	if strings.Contains(str, " ") {
		return strconv.Quote(str)
	}
	return str
}

// Fields collection of information helpful for formatters.
type Fields struct {
	fields []Field
}

// New creates new empty Fields collection.
func New() Fields {
	return Fields{
		fields: []Field{},
	}
}

// List returns fields as an array.
// It'll make a copy of inner fields.
func (f Fields) List() []Field {
	fields := make([]Field, len(f.fields))
	copy(fields, f.fields)
	return fields
}

// Find searches for Field with given name.
// Returns Field's value or nil, and status of existing.
func (f Fields) Find(name string) (interface{}, bool) {
	for i := range f.fields {
		if f.fields[i].Name == name {
			return f.fields[i].Value, true
		}
	}
	return nil, false
}

// Has checks if Field with given name exists.
func (f Fields) Has(name string) bool {
	_, ok := f.Find(name)
	return ok
}

// Get searches for Field with given name and returns
// its value or nil.
func (f Fields) Get(name string) interface{} {
	value, _ := f.Find(name)
	return value
}

// With adds new field or replaces existing.
// It creates a copy of Fields.
func (f Fields) With(name string, value interface{}) Fields {
	for i := range f.fields {
		if f.fields[i].Name == name {
			nf := New()
			nf.fields = f.List()
			nf.fields[i] = Field{
				Name:  name,
				Value: value,
			}
			return nf
		}
	}
	fields := make([]Field, len(f.fields)+1)
	copy(fields, f.fields)
	fields[len(f.fields)] = Field{
		Name:  name,
		Value: value,
	}
	nf := New()
	nf.fields = fields
	return nf
}
