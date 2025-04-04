package mpt

import "fmt"

type LeafNode struct {
	Path  []Nibble
	value any
}

var _ Node = (*LeafNode)(nil)

func NewLeafNodeFromNibbleBytes(nibbles []byte, value any) (*LeafNode, error) {
	ns, err := FromNibbleBytes(nibbles)
	if err != nil {
		return nil, fmt.Errorf("could not leaf node from nibbles: %w", err)
	}

	return NewLeafNodeFromNibbles(ns, value), nil
}

func NewLeafNodeFromNibbles(nibbles []Nibble, value any) *LeafNode {
	return &LeafNode{
		Path:  nibbles,
		value: value,
	}
}

func NewLeafNodeFromKeyValue(key, value string) *LeafNode {
	return NewLeafNodeFromString(key, value)
}

func NewLeafNodeFromBytes(key []byte, value any) *LeafNode {
	return NewLeafNodeFromNibbles(FromBytes(key), value)
}

func NewLeafNodeFromString(key string, value any) *LeafNode {
	return NewLeafNodeFromNibbles(FromString(key), value)
}

func (l *LeafNode) Value() any {
	return l.value
}
