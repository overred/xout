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

// With creates new Fields collection and adds field.
func With(name string, value interface{}) Fields {
	return New().With(name, value)
}

// List returns fields as an array.
// It'll make a copy of inner fields.
func (f Fields) List() []Field {
	fields := make([]Field, len(f.fields))
	copy(fields, f.fields)
	return fields
}

// Count returns number of Fields.
func (f Fields) Count() int {
	return len(f.fields)
}

// Index returns Field by index.
func (f Fields) Index(i int) Field {
	return f.fields[i]
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

// Merge adds given fields to new instance based on current.
// It overwrites fields with same names.
func (f Fields) Merge(fields Fields) Fields {
	nf := New()
	nf.fields = make([]Field, len(f.fields), len(f.fields)+len(fields.fields))
	copy(nf.fields, f.fields)

	for i := range fields.fields {
		for j := range f.fields {
			if fields.fields[i].Name == f.fields[j].Name {
				nf.fields[j] = fields.fields[i]
				goto skip
			}
		}
		nf.fields = append(nf.fields, fields.fields[i])
	skip:
	}

	return nf
}
