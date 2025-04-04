package mpt

import "fmt"

type LeafNode struct {
	Path  []Nibble
	Value []byte
}

var _ Node = (*LeafNode)(nil)

func NewLeafNodeFromNibbleBytes(nibbles []byte, value []byte) (*LeafNode, error) {
	ns, err := FromNibbleBytes(nibbles)
	if err != nil {
		return nil, fmt.Errorf("could not leaf node from nibbles: %w", err)
	}

	return NewLeafNodeFromNibbles(ns, value), nil
}

func NewLeafNodeFromNibbles(nibbles []Nibble, value []byte) *LeafNode {
	return &LeafNode{
		Path:  nibbles,
		Value: value,
	}
}

func NewLeafNodeFromKeyValue(key, value string) *LeafNode {
	return NewLeafNodeFromBytes([]byte(key), []byte(value))
}

func NewLeafNodeFromBytes(key, value []byte) *LeafNode {
	return NewLeafNodeFromNibbles(FromBytes(key), value)
}

func (l *LeafNode) Bytes() []byte {
	return l.Value
}
