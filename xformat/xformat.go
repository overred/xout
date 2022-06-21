package xformat

import (
	"github.com/overred/xout/xfields"
	"github.com/overred/xout/xlevel"
)

// Formatter describes formatter func.
type Formatter func(level xlevel.Level, input string, fields xfields.Fields) string
