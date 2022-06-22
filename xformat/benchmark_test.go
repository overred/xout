package xformat

import (
	"fmt"
	"io"
	"testing"

	"github.com/overred/xout/xfields"
	"github.com/overred/xout/xlevel"
)

// config describes single benchmark configuration for formatter.
type config struct {
	Level  xlevel.Level
	Fields xfields.Fields
}

// matrix describes tests applied for every formatter
var matrix = []config{
	{
		Level:  xlevel.Text,
		Fields: xfields.New(),
	},
	{
		Level:  xlevel.Info,
		Fields: xfields.New(),
	},
	{
		Level: xlevel.Info,
		Fields: xfields.New().
			With("key1", "val1"),
	},
	{
		Level: xlevel.Info,
		Fields: xfields.New().
			With("key1", "val1").
			With("key2", "val2"),
	},
	{
		Level: xlevel.Info,
		Fields: xfields.New().
			With("key1", "val1").
			With("key2", "val2").
			With("key3", "val3"),
	},
}

// bench run matrix tests for given benchmark
func bench(b *testing.B, formatter Formatter) {
	for _, conf := range matrix {
		b.Run(fmt.Sprintf("%s[%d]", conf.Level.String(), len(conf.Fields.List())), func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					formatter.
						Writer(io.Discard, conf.Level, conf.Fields).
						Write([]byte("formatter benchmark text"))
				}
			})
		})
	}
}

func Benchmark_Format_XoutText(b *testing.B) {
	bench(b, NewText())
}

func Benchmark_Format_LogrusText(b *testing.B) {
	bench(b, NewLogrusText())
}

func Benchmark_Format_LogrusJson(b *testing.B) {
	bench(b, NewLogrusJson())
}
