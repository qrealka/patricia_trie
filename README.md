# Patricia trie

Testing patricia trie implementations.

* test some edge cases
* do benchmarking with hashmap

Motivation: find out when better to use patricia trie than hashmap.

## Candidates

* Merkle Patricia Trie
* [Generic Patricia Trie](https://pkg.go.dev/github.com/porfirion/trie#section-readme)

## Banchmark cases

Benchmarks are mostly focused to lookups.

* Random generates short strings (length < 15) and 10k of them.
* Random generates long strings (length > 50) and 10k of them.

```bash
$ go test -bench=.
goos: linux
goarch: amd64
pkg: trie/benchmark
cpu: Intel(R) Xeon(R) E-2176M  CPU @ 2.70GHz
BenchmarkTrieShort-12            7901353             128.3 ns/op
BenchmarkMapShort-12            54117243             19.92 ns/op
BenchmarkTrieLong-12             2386790             490.4 ns/op
BenchmarkMapLong-12             42620047             27.95 ns/op
BenchmarkMapInsertShort-12          2734            442523 ns/op
BenchmarkTrieInsertShort-12          486           2116423 ns/op
BenchmarkMapInsertLong-12           6182            194883 ns/op
BenchmarkTrieInsertLong-12          2278            524100 ns/op
PASS
ok      trie/benchmark  9.305s
```