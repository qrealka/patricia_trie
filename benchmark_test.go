package main

import (
	"fmt"
	"math/rand"
	"testing"

	mpt "merkle-patrica-trie"

	quicktrie "github.com/apokalyptik/quicktrie"
	rTrie "github.com/dghubble/trie"
	pTrie "github.com/porfirion/trie"
)

const (
	numItems       = 10000
	shortStrMinLen = 3
	shortStrMaxLen = 15
	longStrMinLen  = 50
	longStrMaxLen  = 100
	lookupCount    = 1000
)

var (
	shortStrings []string
	longStrings  []string
	lookupShort  []string
	lookupLong   []string
	values       []any
)

func init() {
	// In Go 1.20+, the global random source is automatically seeded
	// No need to manually seed anymore

	// Generate short strings
	shortStrings = make([]string, numItems)
	for i := range numItems {
		shortStrings[i] = generateRandomString(shortStrMinLen, shortStrMaxLen)
	}

	// Generate long strings
	longStrings = make([]string, numItems)
	for i := range numItems {
		longStrings[i] = generateRandomString(longStrMinLen, longStrMaxLen)
	}

	// Generate values (a mix of different types)
	values = make([]any, numItems)
	for i := range numItems {
		switch i % 3 {
		case 0:
			values[i] = fmt.Sprintf("value-%d", i)
		case 1:
			values[i] = i
		case 2:
			values[i] = struct {
				ID   int
				Name string
			}{i, fmt.Sprintf("item-%d", i)}
		}
	}

	// Select lookup keys (75% existing, 25% non-existing)
	lookupShort = make([]string, lookupCount)
	lookupLong = make([]string, lookupCount)

	for i := range lookupCount {
		if i < lookupCount*3/4 {
			// 75% existing keys
			lookupShort[i] = shortStrings[rand.Intn(numItems)]
			lookupLong[i] = longStrings[rand.Intn(numItems)]
		} else {
			// 25% non-existing keys
			lookupShort[i] = generateRandomString(shortStrMinLen, shortStrMaxLen) + "nonexistent"
			lookupLong[i] = generateRandomString(longStrMinLen, longStrMaxLen) + "nonexistent"
		}
	}
}

// generateRandomString creates a random string of length between min and max
func generateRandomString(min, max int) string {
	length := rand.Intn(max-min+1) + min
	bytes := make([]byte, length)
	for i := range length {
		bytes[i] = byte(rand.Intn(26) + 97) // a-z
	}
	return string(bytes)
}

// benchmarkTrieShort tests lookups for MPT with short strings
func BenchmarkTrieShort(b *testing.B) {
	// Create and populate the trie
	trie := mpt.NewTrie()
	for i, key := range shortStrings {
		trie.Put(key, values[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := lookupShort[i%lookupCount]
		_, _ = trie.Get(key)
	}
}

// benchmarkMapShort tests lookups for built-in map with short strings
func BenchmarkMapShort(b *testing.B) {
	// Create and populate the map
	m := make(map[string]any)
	for i, key := range shortStrings {
		m[key] = values[i]
	}

	b.ResetTimer()
	for i := 0; b.Loop(); i++ {
		key := lookupShort[i%lookupCount]
		_ = m[key]
	}
}

// benchmarkTrieLong tests lookups for MPT with long strings
func BenchmarkTrieLong(b *testing.B) {
	// Create and populate the trie
	trie := mpt.NewTrie()
	for i, key := range longStrings {
		trie.Put(key, values[i])
	}

	b.ResetTimer()
	for i := 0; b.Loop(); i++ {
		key := lookupLong[i%lookupCount]
		_, _ = trie.Get(key)
	}
}

// benchmarkMapLong tests lookups for built-in map with long strings
func BenchmarkMapLong(b *testing.B) {
	// Create and populate the map
	m := make(map[string]any)
	for i, key := range longStrings {
		m[key] = values[i]
	}

	b.ResetTimer()
	for i := 0; b.Loop(); i++ {
		key := lookupLong[i%lookupCount]
		_ = m[key]
	}
}

// BenchmarkPTrieShort tests lookups for porfirion/trie with short strings
func BenchmarkPTrieShort(b *testing.B) {
	// Create and populate the trie
	tr := &pTrie.Trie[any]{}
	for i, key := range shortStrings {
		tr.PutString(key, values[i])
	}

	b.ResetTimer()
	for i := 0; b.Loop(); i++ {
		key := lookupShort[i%lookupCount]
		_, _ = tr.GetByString(key)
	}
}

// BenchmarkPTrieLong tests lookups for porfirion/trie with long strings
func BenchmarkPTrieLong(b *testing.B) {
	// Create and populate the trie
	tr := &pTrie.Trie[any]{}
	for i, key := range longStrings {
		tr.PutString(key, values[i])
	}
	b.ResetTimer()
	for i := 0; b.Loop(); i++ {
		key := lookupLong[i%lookupCount]
		_, _ = tr.GetByString(key)
	}
}

// BenchmarkQTrieShort tests lookups for porfirion/trie with short strings
func BenchmarkQTrieShort(b *testing.B) {
	// Create and populate the trie
	tr := quicktrie.NewKVTrie()
	for i, key := range shortStrings {
		tr.Add(key, values[i])
	}

	b.ResetTimer()
	for i := 0; b.Loop(); i++ {
		key := lookupShort[i%lookupCount]
		_, _ = tr.Get(key)
	}
}

// BenchmarkQTrieLong tests lookups for porfirion/trie with long strings
func BenchmarkQTrieLong(b *testing.B) {
	// Create and populate the trie
	tr := quicktrie.NewKVTrie()
	for i, key := range longStrings {
		tr.Add(key, values[i])
	}
	b.ResetTimer()
	for i := 0; b.Loop(); i++ {
		key := lookupLong[i%lookupCount]
		_, _ = tr.Get(key)
	}
}

// BenchmarkRTrieShort tests lookups for porfirion/trie with short strings
func BenchmarkRTrieShort(b *testing.B) {
	// Create and populate the trie
	tr := rTrie.NewRuneTrie()
	for i, key := range shortStrings {
		tr.Put(key, values[i])
	}

	b.ResetTimer()
	for i := 0; b.Loop(); i++ {
		key := lookupShort[i%lookupCount]
		_ = tr.Get(key)
	}
}

// BenchmarkRTrieLong tests lookups for porfirion/trie with long strings
func BenchmarkRTrieLong(b *testing.B) {
	// Create and populate the trie
	tr := rTrie.NewRuneTrie()
	for i, key := range longStrings {
		tr.Put(key, values[i])
	}
	b.ResetTimer()
	for i := 0; b.Loop(); i++ {
		key := lookupLong[i%lookupCount]
		_ = tr.Get(key)
	}
}

// Run with: go test -bench=. -benchmem
// DEFAULT ALLOCATOR
// goos: linux
// goarch: amd64
// pkg: trie/benchmark
// cpu: Intel(R) Xeon(R) E-2176M  CPU @ 2.70GHz
// BenchmarkTrieShort-12            9220322               126.9 ns/op
// BenchmarkMapShort-12            55312198                19.81 ns/op
// BenchmarkTrieLong-12             2371813               496.8 ns/op
// BenchmarkMapLong-12             39583788                29.50 ns/op
// BenchmarkPTrieShort-12          19275516                61.67 ns/op
// BenchmarkPTrieLong-12           10869930               108.0 ns/op

// ARENA ALLOCATOR
