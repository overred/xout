package xformat

import (
	"fmt"
	"io"
	"testing"

	"github.com/overred/xout/xfield"
	"github.com/overred/xout/xlevel"
)

// config describes single benchmark configuration for formatter.
type config struct {
	Level  xlevel.Level
	Fields xfield.Fields
}

// matrix describes tests applied for every formatter
var matrix = []config{
	{
		Level:  xlevel.Text,
		Fields: xfield.New(),
	},
	{
		Level:  xlevel.Info,
		Fields: xfield.New(),
	},
	{
		Level: xlevel.Info,
		Fields: xfield.New().
			With("key1", "val1"),
	},
	{
		Level: xlevel.Info,
		Fields: xfield.New().
			With("key1", "val1").
			With("key2", "val2"),
	},
	{
		Level: xlevel.Info,
		Fields: xfield.New().
			With("key1", "val1").
			With("key2", "val2").
			With("key3", "val3"),
	},
}

// bench run matrix tests for given benchmark
func bench(b *testing.B, formatter Formatter) {
	for _, conf := range matrix {
		// Writer cache for hot writing
		writer := formatter.Writer(io.Discard, conf.Level, conf.Fields)
		b.Run(fmt.Sprintf("%s[%d]", conf.Level.String(), len(conf.Fields.List())), func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					writer.Write([]byte("formatter benchmark text"))
				}
			})
		})
	}
}

func Benchmark_Format_XoutText(b *testing.B) {
	bench(b, NewText())
}

func Benchmark_Format_XoutFastText(b *testing.B) {
	bench(b, NewFastText())
}

// func Benchmark_Format_LogrusText(b *testing.B) {
// 	bench(b, NewLogrusText())
// }

// func Benchmark_Format_LogrusJson(b *testing.B) {
// 	bench(b, NewLogrusJson())
// }
