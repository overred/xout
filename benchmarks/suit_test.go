package benchmarks

import (
	"io"
	"testing"

	"github.com/overred/xout"
	"github.com/sirupsen/logrus"
)

func Benchmark_XOut(b *testing.B) {
	b.Run("DEFA STD", func(b *testing.B) {
		lg := xout.NewPresetDefault()
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				lg.Info("benchmark data")
			}
		})
	})
	b.Run("DEFA FLDS", func(b *testing.B) {
		lg := xout.NewPresetDefault().
			WithField("key1", "val1").
			WithField("key2", "val2").
			WithField("key3", "val3")
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				lg.Info("benchmark data")
			}
		})
	})
	b.Run("DEFA DBG", func(b *testing.B) {
		lg := xout.NewPresetDefault(true)
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				lg.Info("benchmark data")
			}
		})
	})
	b.Run("DEFA DBG FLDS", func(b *testing.B) {
		lg := xout.NewPresetDefault(true).
			WithField("key1", "val1").
			WithField("key2", "val2").
			WithField("key3", "val3")
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				lg.Info("benchmark data")
			}
		})
	})
	b.Run("PERF STD", func(b *testing.B) {
		lg := xout.NewPresetPerformance()
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				lg.Info("benchmark data")
			}
		})
	})
	b.Run("PERF STD FLDS", func(b *testing.B) {
		lg := xout.NewPresetPerformance().
			WithField("key1", "val1").
			WithField("key2", "val2").
			WithField("key3", "val3")
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				lg.Info("benchmark data")
			}
		})
	})
	b.Run("PERF DBG", func(b *testing.B) {
		lg := xout.NewPresetPerformance()
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				lg.Info("benchmark data")
			}
		})
	})
	b.Run("PERF DBG FLDS", func(b *testing.B) {
		lg := xout.NewPresetPerformance().
			WithField("key1", "val1").
			WithField("key2", "val2").
			WithField("key3", "val3")
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				lg.Info("benchmark data")
			}
		})
	})
}

func Benchmark_Logrus(b *testing.B) {
	b.Run("DEFA STD", func(b *testing.B) {
		lg := logrus.New()
		lg.SetOutput(io.Discard)
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				lg.Info("benchmark data")
			}
		})
	})
	b.Run("DEFA STD FLDS", func(b *testing.B) {
		lg := logrus.New()
		lg.SetOutput(io.Discard)
		ent := lg.
			WithField("key1", "val1").
			WithField("key2", "val2").
			WithField("key3", "val3")
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				ent.Info("benchmark data")
			}
		})
	})
	b.Run("DEFA DBG", func(b *testing.B) {
		lg := logrus.New()
		lg.SetOutput(io.Discard)
		lg.SetReportCaller(true)
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				lg.Info("benchmark data")
			}
		})
	})
	b.Run("DEFA DBG FLDS", func(b *testing.B) {
		lg := logrus.New()
		lg.SetOutput(io.Discard)
		lg.SetReportCaller(true)
		ent := lg.
			WithField("key1", "val1").
			WithField("key2", "val2").
			WithField("key3", "val3")
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				ent.Info("benchmark data")
			}
		})
	})
}
