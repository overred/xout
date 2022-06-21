package xfields

// Fields contains specific fields set with info for formatters.
type Fields []Field

// New creates new fields collection.
func New() Fields {
	return make(Fields, 0)
}

// AddField returns new collection with field.
// It overrides field with same name.
func (f Fields) AddField(field Field) Fields {
	fields := make(Fields, len(f), len(f)+1)
	copy(fields, f)
	for i := range fields {
		if fields[i].name == field.name {
			fields[i] = field
			return fields
		}
	}
	return append(fields, field)
}

// AddNameVal returns new collection with field.
// It overrides field with same name.
func (f Fields) AddNameVal(name string, value interface{}) Fields {
	return f.AddField(NewField().WithName(name).WithValue(value))
}
