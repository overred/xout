package xformat

import (
	"io"

	"github.com/overred/xout/xfields"
	"github.com/overred/xout/xlevel"
)

// Formatter describes data Output preprocessor.
type Formatter interface {
	Writer(output io.Writer, level xlevel.Level, fields xfields.Fields) io.Writer
}
