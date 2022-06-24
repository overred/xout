package main

import (
	"log/syslog"
	"os"

	"github.com/overred/xout"
	"github.com/overred/xout/xfields"
	"github.com/overred/xout/xformat"
	"github.com/overred/xout/xlevel"
	"github.com/overred/xout/xposix"
	"github.com/overred/xout/xtarget"
)

func ExampleOverall() {
	slog, _ := syslog.New(syslog.LOG_ERR, "xout")
	xlg := xout.New().
		WithTarget(
			xtarget.New(os.Stdout).
				WithLevelMask(xlevel.AllInfos|xlevel.AllDebugs|xlevel.Text).
				WithPosixMode(xposix.Auto).
				WithFormatter(xformat.NewText()),
			xtarget.New(os.Stderr).
				WithLevelMask(xlevel.AllErrors).
				WithPosixMode(xposix.Auto).
				WithFormatter(xformat.NewText()),
		).
		WithTarget(
			xtarget.Target{
				Output:    slog,
				LevelMask: xlevel.AllErrors,
			},
		).
		WithFields(
			xfields.
				With("a", "1").
				With("b", 5),
			xfields.
				New().
				With("c", nil),
		).
		WithField("d", false).
		WithCaller(true).
		WithTags(true)

	xlg.Println("Hi <fg=red;bg=cyan>there</>!")
	xlg.Trace("Hi <fg=red;bg=cyan>there</>!")
	xlg.Debug("Hi <fg=red;bg=cyan>there!</>!")
	xlg.Info("Hi <fg=red;bg=cyan>there!</>!")
	xlg.Warn("Hi <fg=red;bg=cyan>there!</>!")
	xlg.Error("Hi <fg=red;bg=cyan>there!</>!")
	// xlg.Fatal("Hi <fg=red;bg=cyan>there!</>!") // Will call os.Exit
	// xlg.Panic("Hi <fg=red;bg=cyan>there!</>!") // Will call panic
}
