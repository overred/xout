package xout

import (
	"os"

	"github.com/overred/xout/xformat"
	"github.com/overred/xout/xlevel"
	"github.com/overred/xout/xposix"
	"github.com/overred/xout/xtarget"
)

// NewPresetDefault preset optimized for usability with targets to terminal.
//
// Production:
//  - Targets  : os.Stderr (Error, Fatal, & Panic), & os.Stdout (other).
//  - Levels   : all non-debugging (not Trace, not Debug).
//  - POSIX    : automatic (colored if terminal only).
//  - Tags     : enabled (gookit/color tags).
//  - Caller   : disabled (no info about caller function).
//  - Formatter: XOut Text.
//
// Debug:
//  - Targets  : os.Stderr (Error, Fatal, & Panic), & os.Stdout (other).
//  - Levels   : all (even Trace, & Debug).
//  - POSIX    : automatic (colored if terminal only).
//  - Tags     : enabled (gookit/color tags).
//  - Caller   : enabled (info about caller function).
//  - Formatter: XOut Text.
func NewPresetDefault(debug ...bool) Logger {
	if len(debug) == 0 || !debug[0] {
		return New().
			WithTargets(xtarget.Target{
				Output:    os.Stderr,
				LevelMask: xlevel.AllErrors,
				PosixMode: xposix.Auto,
				Formatter: xformat.NewText(),
			}).
			WithTargets(xtarget.Target{
				Output:    os.Stdout,
				LevelMask: xlevel.AllInfos | xlevel.Text,
				PosixMode: xposix.Auto,
				Formatter: xformat.NewText(),
			}).
			WithTags(true)
	}
	return New().
		WithTargets(xtarget.Target{
			Output:    os.Stderr,
			LevelMask: xlevel.AllErrors,
			PosixMode: xposix.Auto,
			Formatter: xformat.NewText(),
		}).
		WithTargets(xtarget.Target{
			Output:    os.Stdout,
			LevelMask: xlevel.AllInfos | xlevel.AllDebugs | xlevel.Text,
			PosixMode: xposix.Auto,
			Formatter: xformat.NewText(),
		}).
		WithCaller(true).
		WithTags(true)
}

// NewPresetPerformance preset optimized for performance with targets to terminal.
//
// Production:
//  - Targets  : os.Stderr (Error, Fatal, & Panic), & os.Stdout (other).
//  - Levels   : all non-debugging (not Trace, not Debug).
//  - POSIX    : raw (colors pass as is).
//  - Tags     : disabled (gookit/color tags passed as is).
//  - Caller   : disabled (no info about caller function).
//  - Formatter: XOut Fast Text.
//
// Debug:
//  - Targets  : os.Stderr (Error, Fatal, & Panic), & os.Stdout (other).
//  - Levels   : all (even Trace, & Debug).
//  - POSIX    : raw (colors pass as is).
//  - Tags     : disabled (gookit/color tags passed as is).
//  - Caller   : enabled (info about caller function).
//  - Formatter: XOut Fast Text.
func NewPresetPerformance(debug ...bool) Logger {
	if len(debug) == 0 || !debug[0] {
		return New().
			WithTargets(xtarget.Target{
				Output:    os.Stderr,
				LevelMask: xlevel.AllErrors,
				PosixMode: xposix.Raw,
				Formatter: xformat.NewFastText(),
			}).
			WithTargets(xtarget.Target{
				Output:    os.Stdout,
				LevelMask: xlevel.AllInfos | xlevel.Text,
				PosixMode: xposix.Raw,
				Formatter: xformat.NewFastText(),
			})
	}
	return New().
		WithTargets(xtarget.Target{
			Output:    os.Stderr,
			LevelMask: xlevel.AllErrors,
			PosixMode: xposix.Raw,
			Formatter: xformat.NewFastText(),
		}).
		WithTargets(xtarget.Target{
			Output:    os.Stdout,
			LevelMask: xlevel.AllInfos | xlevel.AllDebugs | xlevel.Text,
			PosixMode: xposix.Raw,
			Formatter: xformat.NewFastText(),
		}).
		WithCaller(true)
}
