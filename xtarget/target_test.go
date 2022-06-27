package xtarget

import (
	"io"
	"testing"

	"github.com/overred/xout/xfield"
	"github.com/overred/xout/xlevel"
	"github.com/overred/xout/xposix"
)

// Prepared fields set.
var fields = xfield.New()

// The minimal and lightweight target.
var targetPosixRaw = Target{
	Output:    io.Discard,
	LevelMask: xlevel.All,
	PosixMode: xposix.Raw,
	Formatter: nil,
}

// The target with force POSIX clean mode.
var targetPosixClear = Target{
	Output:    io.Discard,
	LevelMask: xlevel.All,
	PosixMode: xposix.Clean,
	Formatter: nil,
}

// The target with automatic POSIX.
var targetPosixAuto = Target{
	Output:    io.Discard,
	LevelMask: xlevel.All,
	PosixMode: xposix.Auto,
	Formatter: nil,
}

func Benchmark_Target_PosixRaw(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = targetPosixRaw.Writer(xlevel.Debug, fields)
		}
	})
}

func Benchmark_Target_PosixClear(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = targetPosixClear.Writer(xlevel.Debug, fields)
		}
	})
}

func Benchmark_Target_PosixAuto(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = targetPosixAuto.Writer(xlevel.Debug, fields)
		}
	})
}
