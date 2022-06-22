package main

import (
	"os"

	"github.com/overred/xout/xfields"
	"github.com/overred/xout/xformat"
	"github.com/overred/xout/xlevel"
	"github.com/overred/xout/xtarget"
)

func main() {
	xtarget.New(os.Stdout).
		WithFormatter(xformat.NewText()).
		Writer(xlevel.Info, xfields.New().With("logger", "xout")).
		Write([]byte("Hi there!"))
}
