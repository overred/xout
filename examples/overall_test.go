package main

import (
	"errors"
	"log/syslog"
	"os"

	"github.com/overred/xout"
	"github.com/overred/xout/xfield"
	"github.com/overred/xout/xformat"
	"github.com/overred/xout/xlevel"
	"github.com/overred/xout/xposix"
	"github.com/overred/xout/xtarget"
)

func ExampleOverall() {
	// Syslog connection
	slog, _ := syslog.New(syslog.LOG_ERR, "xout")

	// Logger with custom config
	xlg := xout.New().
		WithTargets(
			xtarget.New(os.Stdout).
				WithLevelMask(xlevel.AllInfos|xlevel.AllDebugs|xlevel.Text).
				WithPosixMode(xposix.Auto).
				WithFormatter(xformat.NewText()),
			xtarget.New(os.Stderr).
				WithLevelMask(xlevel.AllErrors).
				WithPosixMode(xposix.Auto).
				WithFormatter(xformat.NewText()),
		).
		WithTargets(
			xtarget.Target{
				Output:    slog,
				LevelMask: xlevel.AllErrors,
			},
		).
		WithFields(
			xfield.
				With("a", "1").
				With("b", 5),
			xfield.
				New().
				With("c", nil),
		).
		WithField("d", false).
		WithError(errors.New("err example")).
		WithCaller(true).
		WithTags(true)

	xlg.Println("Hi <fg=lightRed;bg=cyan>there</>!")
	xlg.Trace("Hi <fg=lightRed;bg=cyan>there</>!")
	xlg.Debug("Hi <fg=lightRed;bg=cyan>there!</>!")
	xlg.Info("Hi <fg=lightRed;bg=cyan>there!</>!")
	xlg.Warn("Hi <fg=lightRed;bg=cyan>there!</>!")
	xlg.Error("Hi <fg=lightRed;bg=cyan>there!</>!")
	xlg.Fatal("Hi <fg=lightRed;bg=cyan>there!</>!") // Will call os.Exit
	xlg.Panic("Hi <fg=LightRed;bg=cyan>there!</>!") // Will call panic
}
