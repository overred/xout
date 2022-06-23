package xout

import (
	"os"

	"github.com/overred/xout/xformat"
	"github.com/overred/xout/xlevel"
	"github.com/overred/xout/xposix"
	"github.com/overred/xout/xtarget"
)

// NewPresetDefault preset optimized for usability.
//
//  - os.Stdin & os.Stdout as the targets
//  - all non debug levels are allowed
//  - automatic POSIX mode
//  - color tags enabled
//  - default text formatter
func NewPresetDefault() Logger {
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

// NewPresetDebugDefault preset optimized for usability
// and debugging.
//
//  - os.Stdin & os.Stdout as the targets
//  - all levels are allowed
//  - automatic POSIX mode
//  - color tags enabled
//  - default text formatter
func NewPresetDebugDefault() Logger {
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
		WithTags(true)
}

// NewPresetFormat preset optimized for usability with custom formatter.
//
//  - os.Stdin & os.Stdout as the targets
//  - all non debug levels are allowed
//  - automatic POSIX mode
//  - color tags enabled
//  - formatter by Your choice
func NewPresetFormat(formatter xformat.Formatter) Logger {
	return New().
		WithTarget(xtarget.Target{
			Output:    os.Stderr,
			LevelMask: xlevel.AllErrors,
			PosixMode: xposix.Auto,
			Formatter: formatter,
		}).
		WithTarget(xtarget.Target{
			Output:    os.Stdout,
			LevelMask: xlevel.AllInfos | xlevel.Text,
			PosixMode: xposix.Auto,
			Formatter: formatter,
		}).
		WithTags(true)
}

// NewPresetDebugFormat preset optimized for usability with custom formatter
// and debugging.
//
//  - os.Stdin & os.Stdout as the targets
//  - all levels are allowed
//  - automatic POSIX mode
//  - color tags enabled
//  - formatter by Your choice
func NewPresetDebugFormat(formatter xformat.Formatter) Logger {
	return New().
		WithTarget(xtarget.Target{
			Output:    os.Stderr,
			LevelMask: xlevel.AllErrors,
			PosixMode: xposix.Auto,
			Formatter: formatter,
		}).
		WithTarget(xtarget.Target{
			Output:    os.Stdout,
			LevelMask: xlevel.AllInfos | xlevel.AllDebugs | xlevel.Text,
			PosixMode: xposix.Auto,
			Formatter: formatter,
		}).
		WithTags(true)
}

// NewPresetFastText preset optimized for performance.
//
//  - os.Stdin & os.Stdout as the targets
//  - all non debug levels are allowed
//  - raw POSIX mode
//  - color tags disabled
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
		})
}
