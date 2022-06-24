package xout

import (
	"os"

	"github.com/overred/xout/xformat"
	"github.com/overred/xout/xlevel"
	"github.com/overred/xout/xposix"
	"github.com/overred/xout/xtarget"
)

// NewPresetText preset optimized for usability.
//
//  - os.Stdin & os.Stdout as the targets
//  - all non debug levels are allowed
//  - automatic POSIX mode
//  - color tags enabled
//  - caller fields disabled
//  - default text formatter
func NewPresetText() Logger {
	return New().
		WithTarget(xtarget.Target{
			Output:    os.Stderr,
			LevelMask: xlevel.AllErrors,
			PosixMode: xposix.Auto,
			Formatter: xformat.NewText(),
		}).
		WithTarget(xtarget.Target{
			Output:    os.Stdout,
			LevelMask: xlevel.AllInfos | xlevel.Text,
			PosixMode: xposix.Auto,
			Formatter: xformat.NewText(),
		}).
		WithTags(true)
}

// NewPresetDebugText preset optimized for usability
// and debugging.
//
//  - os.Stdin & os.Stdout as the targets
//  - all levels are allowed
//  - automatic POSIX mode
//  - color tags enabled
//  - caller fields enabled
//  - default text formatter
func NewPresetDebugText() Logger {
	return New().
		WithTarget(xtarget.Target{
			Output:    os.Stderr,
			LevelMask: xlevel.AllErrors,
			PosixMode: xposix.Auto,
			Formatter: xformat.NewText(),
		}).
		WithTarget(xtarget.Target{
			Output:    os.Stdout,
			LevelMask: xlevel.AllInfos | xlevel.AllDebugs | xlevel.Text,
			PosixMode: xposix.Auto,
			Formatter: xformat.NewText(),
		}).
		WithCaller(true).
		WithTags(true)
}

// NewPresetFastText preset optimized for performance.
//
//  - os.Stdin & os.Stdout as the targets
//  - all non debug levels are allowed
//  - raw POSIX mode
//  - color tags disabled
//  - caller fields disabled
//  - fastest text formatter
func NewPresetFastText() Logger {
	return New().
		WithTarget(xtarget.Target{
			Output:    os.Stderr,
			LevelMask: xlevel.AllErrors,
			PosixMode: xposix.Raw,
			Formatter: xformat.NewFastText(),
		}).
		WithTarget(xtarget.Target{
			Output:    os.Stdout,
			LevelMask: xlevel.AllInfos | xlevel.Text,
			PosixMode: xposix.Raw,
			Formatter: xformat.NewFastText(),
		})
}

// NewPresetDebugFastText preset optimized for performance
// and debugging.
//
//  - os.Stdin & os.Stdout as the targets
//  - all levels are allowed
//  - raw POSIX mode
//  - color tags disabled
//  - caller fields enabled
//  - fastest text formatter
func NewPresetDebugFastText() Logger {
	return New().
		WithTarget(xtarget.Target{
			Output:    os.Stderr,
			LevelMask: xlevel.AllErrors,
			PosixMode: xposix.Raw,
			Formatter: xformat.NewFastText(),
		}).
		WithTarget(xtarget.Target{
			Output:    os.Stdout,
			LevelMask: xlevel.AllInfos | xlevel.AllDebugs | xlevel.Text,
			PosixMode: xposix.Raw,
			Formatter: xformat.NewFastText(),
		}).
		WithCaller(true)
}
