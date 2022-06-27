package xout

import (
	"fmt"
	"testing"

	"github.com/overred/xout/xfield"
)

func Benchmark_Logger(b *testing.B) {
	var matrix = []struct {
		Name   string
		Logger Logger
	}{
		{
			Name:   "Defa Prod",
			Logger: NewPresetDefault(),
		},
		{
			Name:   "Defa Debg",
			Logger: NewPresetDefault(true),
		},
		{
			Name:   "Perf Prod",
			Logger: NewPresetPerformance(),
		},
		{
			Name:   "Perf Debg",
			Logger: NewPresetPerformance(true),
		},
	}
	for _, bench := range matrix {
		b.Run(bench.Name+"/Info", func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					bench.Logger.Info("presets benchmark")
				}
			})
		})
		b.Run(bench.Name+"/Flds", func(b *testing.B) {
			fields := xfield.New()
			for i := 0; i < 10; i++ {
				fields = fields.With(fmt.Sprintf("field%d", i), fmt.Sprintf("value%d", i))
			}
			bench.Logger = bench.Logger.WithFields(fields)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					bench.Logger.Info("presets benchmark")
				}
			})
		})
		b.Run(bench.Name+"/Text", func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					bench.Logger.Print("presets benchmark")
				}
			})
		})
	}
}
