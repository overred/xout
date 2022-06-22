package xtarget

import (
	"io"
	"testing"

	"github.com/overred/xout/xfields"
	"github.com/overred/xout/xlevel"
	"github.com/overred/xout/xposix"
)

// The minimal and lightweight writer.
var writerPosixRaw = Target{
	Output:    io.Discard,
	LevelMask: xlevel.All,
	PosixMode: xposix.Raw,
	Formatter: nil,
}.Writer(xlevel.Debug, xfields.New())

// The writer with force POSIX clean mode.
var writerPosixClear = Target{
	Output:    io.Discard,
	LevelMask: xlevel.All,
	PosixMode: xposix.Clear,
	Formatter: nil,
}.Writer(xlevel.Debug, xfields.New())

// The writer with automatic POSIX.
var writerPosixAuto = Target{
	Output:    io.Discard,
	LevelMask: xlevel.All,
	PosixMode: xposix.Auto,
	Formatter: nil,
}.Writer(xlevel.Debug, xfields.New())

func Benchmark_Writer_PosixRaw(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			writerPosixRaw.Write([]byte("posix raw"))
		}
	})
}

func Benchmark_Writer_PosixClear(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			writerPosixClear.Write([]byte("posix clear"))
		}
	})
}

func Benchmark_Writer_PosixAuto(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			writerPosixAuto.Write([]byte("posix auto"))
		}
	})
}
