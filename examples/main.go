package main

import (
	"os"

	"github.com/overred/xout/xfields"
	"github.com/overred/xout/xformat"
	"github.com/overred/xout/xlevel"
	"github.com/overred/xout/xtarget"
)

func main() {
	// o := xout.Out{
	// 	Target: os.Stdout,
	// 	Color:  xout.ColorAuto,
	// 	Level:  xout.LevelInfo.Elevate(),
	// 	Format: xout.NewLogrusTextFormatter(),
	// }
	// o.Write(xout.LevelInfo, []byte("123678 dsads"), map[string]interface{}{"test": 5})
	// o.Write(xout.LevelInfo, []byte(fmt.Sprintf("Test %s", aurora.Red("RED"))), map[string]interface{}{"test": 5})
	// o.Write(xout.LevelWarn, []byte(fmt.Sprintf("Test %s text", aurora.Red("RED"))), map[string]interface{}{"test": 5, "data": `str" spc`})
	// o.Write(xout.LevelWarn, []byte("very very longest text in the universe to test the logger"), map[string]interface{}{"test": 5, "data": `str" spc`})
	// xout.NewOutStdPrinter().Write(xout.LevelText, []byte("text\n"), nil)
	// target.New().WithWriter(os.Stdout).WithFormat().WithLevel(level.Info.Elevate()).Write(level.Info, []byte("test\n"))
	// xtarget.NewStdout().AddNameVal("test", `val "data`).Write(xlevel.Info, []byte("test"))

	// xout.NewDefault().Write(xlevel.Text, "hello <fg=brightred>colored</> regular\n")
	// xout.NewDefault().Write(xlevel.Trace, "hello <fg=lightRed>colored</>")
	// xout.NewDefault().Write(xlevel.Debug, "hello")
	// xout.NewDefault().Write(xlevel.Info, "hello")
	// xout.NewDefault().Write(xlevel.Warn, "hello")
	// xout.NewDefault().Write(xlevel.Error, "hello")
	// xout.NewDefault().Write(xlevel.Fatal, "hello")
	// xout.NewDefault().Write(xlevel.Panic, "hello")
	// xout.NewDefault().Printf("[%s]\n", "data")

	// xout.NewDefault().Fatalf("yaya! <fg=red>PANIC!</>")

	f := xfields.New().With("key", "val 1").With("data", 5)

	x := xtarget.New(os.Stdout)
	xa := x.WithFormatter(xformat.NewDefault())
	xb := x.WithFormatter(xformat.NewLogrusText())

	wa := xa.Writer(xlevel.Info, f)
	wb := xb.Writer(xlevel.Info, f)

	wa.Write([]byte("test\n"))
	wb.Write([]byte("test\n"))
}
