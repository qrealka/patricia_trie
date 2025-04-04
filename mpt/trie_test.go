package mpt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// check basic key-value mapping
func TestGetPut(t *testing.T) {
	t.Run("should get nothing if key does not exist", func(t *testing.T) {
		trie := NewTrie()
		_, found := trie.Get("notexist")
		require.Equal(t, false, found)
	})

	t.Run("should get value if key exist", func(t *testing.T) {
		trie := NewTrie()
		trie.Put("test", "hello")
		val, found := trie.Get("test")
		require.Equal(t, true, found)
		require.Equal(t, "hello", val)
	})

	t.Run("should get updated value", func(t *testing.T) {
		trie := NewTrie()
		trie.Put("test", "hello")
		trie.Put("test", "world")
		val, found := trie.Get("test")
		require.Equal(t, true, found)
		require.Equal(t, "world", val)
	})

	t.Run("should work with any value type", func(t *testing.T) {
		trie := NewTrie()
		trie.Put("number", 42)
		trie.Put("boolean", true)
		trie.Put("struct", struct{ Name string }{"Test"})

		num, found := trie.Get("number")
		require.True(t, found)
		require.Equal(t, 42, num)

		b, found := trie.Get("boolean")
		require.True(t, found)
		require.Equal(t, true, b)

		s, found := trie.Get("struct")
		require.True(t, found)
		require.Equal(t, struct{ Name string }{"Test"}, s)
	})
}

func TestPut2Pairs(t *testing.T) {
	trie := NewTrie()
	trie.Put("verb", "verb")
	trie.Put("verbal", "coin")

	verb, ok := trie.Get("verb")
	require.True(t, ok)
	require.Equal(t, "verb", verb)

	coin, ok := trie.Get("verbal")
	require.True(t, ok)
	require.Equal(t, "coin", coin)
}

func TestPutLeafShorter(t *testing.T) {
	trie := NewTrie()
	trie.Put("test", "hello")
	trie.Put("tes", "world")

	hello, ok := trie.Get("test")
	require.True(t, ok)
	require.Equal(t, "hello", hello)

	world, ok := trie.Get("tes")
	require.True(t, ok)
	require.Equal(t, "world", world)
}
