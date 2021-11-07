1. [How to generate a random string of a fixed length in Go?](https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go)
2. [鸟窝 快速产生一个随机字符串](https://colobu.com/2018/09/02/generate-random-string-in-Go/)

```bash
🕙[2021-11-07 12:31:25.044] ❯ go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/bingoohuang/gg-rand/str
cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
BenchmarkRunes-12                                        2796652               433.0 ns/op            88 B/op          2 allocs/op
BenchmarkBytes-12                                        3688272               315.9 ns/op            32 B/op          2 allocs/op
BenchmarkBytesRmndr-12                                   4733895               256.6 ns/op            32 B/op          2 allocs/op
BenchmarkBytesMask-12                                    4018478               295.3 ns/op            32 B/op          2 allocs/op
BenchmarkBytesMaskImpr-12                               10661510               110.6 ns/op            32 B/op          2 allocs/op
BenchmarkBytesMaskImprSrc-12                            11714509               102.9 ns/op            32 B/op          2 allocs/op
BenchmarkBytesMaskImprXorshift1024Src-12                11524878               102.2 ns/op            32 B/op          2 allocs/op
BenchmarkBytesMaskImprXorshift256Src-12                 12077193                98.86 ns/op           32 B/op          2 allocs/op
BenchmarkBytesMaskImprXorShift64StarSrc-12               9288363               132.2 ns/op            48 B/op          3 allocs/op
BenchmarkBytesMaskImprXorShift128PlusSrc-12              9257799               125.7 ns/op            48 B/op          3 allocs/op
BenchmarkBytesMaskImprXorShift1024StarSrc-12             8942186               133.8 ns/op            48 B/op          3 allocs/op
BenchmarkSecureRandomAlphaString-12                       922449              1356 ns/op              64 B/op          3 allocs/op
BenchmarkSecureRandomString-12                            901616              1248 ns/op              61 B/op          3 allocs/op
BenchmarkShortID-12                                      1445821               824.8 ns/op            32 B/op          2 allocs/op
BenchmarkGenerate-12                                    20054028                58.69 ns/op           16 B/op          1 allocs/op
BenchmarkRandStr-12                                     22999382                49.15 ns/op           16 B/op          1 allocs/op
BenchmarkUniuriNewLen-12                                 1400449               857.6 ns/op            56 B/op          3 allocs/op
PASS
ok      github.com/bingoohuang/gg-rand/str      24.215s

```
