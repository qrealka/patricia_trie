package mpt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// check basic key-value mapping
func TestGetPut(t *testing.T) {
	t.Run("should get nothing if key does not exist", func(t *testing.T) {
		trie := NewTrie()
		_, found := trie.Get([]byte("notexist"))
		require.Equal(t, false, found)
	})

	t.Run("should get value if key exist", func(t *testing.T) {
		trie := NewTrie()
		trie.Put([]byte{1, 2, 3, 4}, []byte("hello"))
		val, found := trie.Get([]byte{1, 2, 3, 4})
		require.Equal(t, true, found)
		require.Equal(t, val, []byte("hello"))
	})

	t.Run("should get updated value", func(t *testing.T) {
		trie := NewTrie()
		trie.Put([]byte{1, 2, 3, 4}, []byte("hello"))
		trie.Put([]byte{1, 2, 3, 4}, []byte("world"))
		val, found := trie.Get([]byte{1, 2, 3, 4})
		require.Equal(t, true, found)
		require.Equal(t, val, []byte("world"))
	})
}

func TestPut2Pairs(t *testing.T) {
	trie := NewTrie()
	trie.Put([]byte{1, 2, 3, 4}, []byte("verb"))
	trie.Put([]byte{1, 2, 3, 4, 5, 6}, []byte("coin"))

	verb, ok := trie.Get([]byte{1, 2, 3, 4})
	require.True(t, ok)
	require.Equal(t, []byte("verb"), verb)

	coin, ok := trie.Get([]byte{1, 2, 3, 4, 5, 6})
	require.True(t, ok)
	require.Equal(t, []byte("coin"), coin)

	fmt.Printf("%T\n", trie.root)
	ext, ok := trie.root.(*ExtensionNode)
	require.True(t, ok)
	branch, ok := ext.Next.(*BranchNode)
	require.True(t, ok)
	_, ok = branch.Branches[0].(*LeafNode)
	require.True(t, ok)
}

func TestPutLeafShorter(t *testing.T) {
	trie := NewTrie()
	trie.Put([]byte{1, 2, 3, 4}, []byte("hello"))
	trie.Put([]byte{1, 2, 3}, []byte("world"))

	leaf := NewLeafNodeFromNibbles([]Nibble{4}, []byte("hello"))

	branch := NewBranchNode()
	branch.SetBranch(Nibble(0), leaf)
	branch.SetValue([]byte("world"))

	_ = NewExtensionNode([]Nibble{0, 1, 0, 2, 0, 3}, branch)
}

// Before put:
//
//  	           ┌───────────────────────────┐
//  	           │  Extension Node           │
//  	           │  Path: [0, 1, 0, 2, 0, 3] │
//  	           └────────────┬──────────────┘
//  	                        │
//  	┌───────────────────────┴──────────────────┐
//  	│                   Branch Node            │
//  	│   [0]         ...          [5]           │
//  	└────┼────────────────────────┼────────────┘
//  	     │                        │
//  	     │                        │
//  	     │                        │
//  	     │                        │
//   ┌───────┴──────────┐   ┌─────────┴─────────┐
//   │  Leaf Node       │   │  Leaf Node        │
//   │  Path: [4]       │   │  Path: [0]        │
//   │  Value: "hello1" │   │  Value: "hello2"  │
//   └──────────────────┘   └───────────────────┘
//
// After put([]byte{[1, 2, 3]}, "world"):
//  	           ┌───────────────────────────┐
//  	           │  Extension Node           │
//  	           │  Path: [0, 1, 0, 2, 0, 3] │
//  	           └────────────┬──────────────┘
//  	                        │
//  	┌───────────────────────┴────────────────────────┐
//  	│                   Branch Node                  │
//  	│   [0]         ...          [5]  value: "world" │
//  	└────┼────────────────────────┼──────────────────┘
//  	     │                        │
//  	     │                        │
//  	     │                        │
//  	     │                        │
//   ┌───────┴──────────┐   ┌─────────┴─────────┐
//   │  Leaf Node       │   │  Leaf Node        │
//   │  Path: [4]       │   │  Path: [0]        │
//   │  Value: "hello1" │   │  Value: "hello2"  │
//   └──────────────────┘   └───────────────────┘
