Pack fractional bytes (`uint4`, `uint6`, `uint7`) into bytes.

![fuzzing](https://img.shields.io/badge/fuzzing-active-brightgreen)



```bash
$ go test -bench=. -benchmem .
goos: darwin
goarch: arm64
pkg: github.com/nikolaydubina/go-bit-packing
Benchmark_2x4b/pack-16         	1000000000	         0.8389 ns/op	         2.220 GB/s	       0 B/op	       0 allocs/op
Benchmark_2x4b/unpack-16       	1000000000	         0.6956 ns/op	         1.339 GB/s	       0 B/op	       0 allocs/op
Benchmark_4x6b/pack-16         	1000000000	         1.047 ns/op	         3.559 GB/s	       0 B/op	       0 allocs/op
Benchmark_4x6b/unpack-16       	453752890	         2.563 ns/op	         1.090 GB/s	       0 B/op	       0 allocs/op
Benchmark_8x7b/pack-16         	399813974	         2.951 ns/op	         2.525 GB/s	       0 B/op	       0 allocs/op
Benchmark_8x7b/unpack-16       	251525857	         4.805 ns/op	         1.357 GB/s	       0 B/op	       0 allocs/op
```
