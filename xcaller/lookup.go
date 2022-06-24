package xcaller

import (
	"runtime"

	"github.com/overred/xout/xfields"
)

const (
	// FieldNameFile contains name for xfields.Fields to get caller's file name.
	FieldNameFile = "caller.file"
	// FieldNameFunc contains name for xfields.Fields to get caller's function name.
	FieldNameFunc = "caller.func"
	// FieldNameLine contains name for xfields.Fields to get caller's line number.
	FieldNameLine = "caller.line"
)

// Lookup searches caller info inside frames.
func Lookup(skip int) (filename string, function string, line int) {
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		fn := runtime.FuncForPC(pc)
		return file, fn.Name(), line
	}
	return "", "", 0
}

// AsFields returns caller info wrapped into xfields.Fields
// to pass it into xformat.Formatter.
func AsFields(skip int) xfields.Fields {
	file, function, line := Lookup(skip)
	fields := xfields.New().
		With(FieldNameFile, file).
		With(FieldNameFunc, function).
		With(FieldNameLine, line)
	return fields
}
