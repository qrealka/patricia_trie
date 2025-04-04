package mpt

import "fmt"

// TrieInterface defines the interface for a Patricia Trie
type TrieInterface interface {
	Get(key string) (any, bool)
	Put(key string, value any)
	Del(key string) bool
}

// Ensure Trie implements TrieInterface
var _ TrieInterface = (*Trie)(nil)

type Trie struct {
	root Node
}

func NewTrie() *Trie {
	return &Trie{}
}

// Get retrieves a value from the trie using a string key
func (t *Trie) Get(key string) (any, bool) {
	node := t.root
	nibbles := FromString(key)
	for {
		if IsEmptyNode(node) {
			return nil, false
		}

		if leaf, ok := node.(*LeafNode); ok {
			matched := PrefixMatchedLen(leaf.Path, nibbles)
			if matched != len(leaf.Path) || matched != len(nibbles) {
				return nil, false
			}
			return leaf.Value(), true
		}

		if branch, ok := node.(*BranchNode); ok {
			if len(nibbles) == 0 {
				return branch.Value(), branch.HasValue()
			}

			b, remaining := nibbles[0], nibbles[1:]
			nibbles = remaining
			node = branch.Branches[b]
			continue
		}

		if ext, ok := node.(*ExtensionNode); ok {
			matched := PrefixMatchedLen(ext.Path, nibbles)
			// E 01020304
			//   010203
			if matched < len(ext.Path) {
				return nil, false
			}

			nibbles = nibbles[matched:]
			node = ext.Next
			continue
		}

		panic("not found")
	}
}

// Put adds a key-value pair to the trie using a string key and any value
func (t *Trie) Put(key string, value any) {
	node := &t.root
	nibbles := FromString(key)
	for {
		if IsEmptyNode(*node) {
			leaf := NewLeafNodeFromNibbles(nibbles, value)
			*node = leaf
			return
		}

		if leaf, ok := (*node).(*LeafNode); ok {
			matched := PrefixMatchedLen(leaf.Path, nibbles)

			// if all matched, update value even if the value are equal
			if matched == len(nibbles) && matched == len(leaf.Path) {
				newLeaf := NewLeafNodeFromNibbles(leaf.Path, value)
				*node = newLeaf
				return
			}

			branch := NewBranchNode()
			// if matched some nibbles, check if matches either all remaining nibbles
			// or all leaf nibbles
			if matched == len(leaf.Path) {
				branch.SetValue(leaf.Value())
			}

			if matched == len(nibbles) {
				branch.SetValue(value)
			}

			// if there are matched nibbles, an extension node will be created
			if matched > 0 {
				// create an extension node for the shared nibbles
				ext := NewExtensionNode(leaf.Path[:matched], branch)
				*node = ext
			} else {
				// when there no matched nibble, there is no need to keep the extension node
				*node = branch
			}

			if matched < len(leaf.Path) {
				// have dismatched
				// L 01020304 hello
				// + 010203   world

				// 01020304, 0, 4
				branchNibble, leafNibbles := leaf.Path[matched], leaf.Path[matched+1:]
				newLeaf := NewLeafNodeFromNibbles(leafNibbles, leaf.Value()) // not :matched+1
				branch.SetBranch(branchNibble, newLeaf)
			}

			if matched < len(nibbles) {
				// L 01020304 hello
				// + 010203040 world

				// L 01020304 hello
				// + 010203040506 world
				branchNibble, leafNibbles := nibbles[matched], nibbles[matched+1:]
				newLeaf := NewLeafNodeFromNibbles(leafNibbles, value)
				branch.SetBranch(branchNibble, newLeaf)
			}

			return
		}

		if branch, ok := (*node).(*BranchNode); ok {
			if len(nibbles) == 0 {
				branch.SetValue(value)
				return
			}

			b, remaining := nibbles[0], nibbles[1:]
			nibbles = remaining
			node = &branch.Branches[b]
			continue
		}

		// E 01020304
		// B 0 hello
		// L 506 world
		// + 010203 good
		if ext, ok := (*node).(*ExtensionNode); ok {
			matched := PrefixMatchedLen(ext.Path, nibbles)
			if matched < len(ext.Path) {
				// E 01020304
				// + 010203 good
				extNibbles, branchNibble, extRemainingnibbles := ext.Path[:matched], ext.Path[matched], ext.Path[matched+1:]
				branch := NewBranchNode()
				if len(extRemainingnibbles) == 0 {
					// E 0102030
					// + 010203 good
					branch.SetBranch(branchNibble, ext.Next)
				} else {
					// E 01020304
					// + 010203 good
					newExt := NewExtensionNode(extRemainingnibbles, ext.Next)
					branch.SetBranch(branchNibble, newExt)
				}

				if matched < len(nibbles) {
					nodeBranchNibble, nodeLeafNibbles := nibbles[matched], nibbles[matched+1:]
					remainingLeaf := NewLeafNodeFromNibbles(nodeLeafNibbles, value)
					branch.SetBranch(nodeBranchNibble, remainingLeaf)
				} else if matched == len(nibbles) {
					branch.SetValue(value)
				} else {
					panic(fmt.Sprintf("too many matched (%v > %v)", matched, len(nibbles)))
				}

				// if there is no shared extension nibbles any more, then we don't need the extension node
				// any more
				// E 01020304
				// + 1234 good
				if len(extNibbles) == 0 {
					*node = branch
				} else {
					// otherwise create a new extension node
					*node = NewExtensionNode(extNibbles, branch)
				}
				return
			}

			nibbles = nibbles[matched:]
			node = &ext.Next
			continue
		}

		panic("unknown type")
	}
}

// Del deletes a key-value pair from the trie using a string key
// Currently a stub that needs implementation
func (t *Trie) Del(key string) bool {
	// TODO: Implement deletion
	return false
}
