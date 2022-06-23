```json
{
    "tags": true,
    "outs": [
        {
            "colors": "raw|no|auto",
            "output": "io.Writer.(*os.File)",
            "levels": "Trace|Debug|Info|Warning|Error|Fatal|Panic|Text",
            "format": "logrus",
        }
    ]
}
```

# Tests
Run all tests
```bash
go test ./...
```

```
go test --cpuprofile=prof_xformat_cpu.prof --memprofile=prof_xformat_mem.prof --bench=. --benchtime=1s --benchmem ./xformat
goos: darwin
goarch: amd64
pkg: github.com/overred/xout/xformat
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
Benchmark_Format_XoutText/text[0]-12            83777868                15.40 ns/op           24 B/op          1 allocs/op
Benchmark_Format_XoutText/info[0]-12             2896276               438.1 ns/op           368 B/op         13 allocs/op
Benchmark_Format_XoutText/info[1]-12             2659352               454.0 ns/op           432 B/op         13 allocs/op
Benchmark_Format_XoutText/info[2]-12             2507262               465.4 ns/op           496 B/op         13 allocs/op
Benchmark_Format_XoutText/info[3]-12             2487963               449.1 ns/op           528 B/op         13 allocs/op
Benchmark_Format_XoutFastText/text[0]-12        79628425                13.32 ns/op           24 B/op          1 allocs/op
Benchmark_Format_XoutFastText/info[0]-12        36012008                27.93 ns/op           56 B/op          2 allocs/op
Benchmark_Format_XoutFastText/info[1]-12        35696890                30.06 ns/op           72 B/op          2 allocs/op
Benchmark_Format_XoutFastText/info[2]-12        35598663                29.91 ns/op           88 B/op          2 allocs/op
Benchmark_Format_XoutFastText/info[3]-12        32385924                32.18 ns/op           88 B/op          2 allocs/op
Benchmark_Format_LogrusText/text[0]-12          63082308                16.22 ns/op           24 B/op          1 allocs/op
Benchmark_Format_LogrusText/info[0]-12           7058965               166.9 ns/op           224 B/op          5 allocs/op
Benchmark_Format_LogrusText/info[1]-12           6455835               168.4 ns/op           256 B/op          5 allocs/op
Benchmark_Format_LogrusText/info[2]-12           6181770               188.2 ns/op           288 B/op          5 allocs/op
Benchmark_Format_LogrusText/info[3]-12           5667346               200.7 ns/op           352 B/op          5 allocs/op
Benchmark_Format_LogrusJson/text[0]-12          70321220                14.52 ns/op           24 B/op          1 allocs/op
Benchmark_Format_LogrusJson/info[0]-12           1174612               937.7 ns/op          1417 B/op         24 allocs/op
Benchmark_Format_LogrusJson/info[1]-12            876216              1281 ns/op            1609 B/op         27 allocs/op
Benchmark_Format_LogrusJson/info[2]-12            741801              1649 ns/op            1769 B/op         29 allocs/op
Benchmark_Format_LogrusJson/info[3]-12            741614              1672 ns/op            1929 B/op         31 allocs/op
PASS
ok      github.com/overred/xout/xformat 29.774s
```