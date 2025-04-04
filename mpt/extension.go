package mpt

type ExtensionNode struct {
	Path []Nibble
	Next Node
}

var _ Node = (*ExtensionNode)(nil)

func NewExtensionNode(nibbles []Nibble, next Node) *ExtensionNode {
	return &ExtensionNode{
		Path: nibbles,
		Next: next,
	}
}

func (e *ExtensionNode) Bytes() []byte {
	return nil
}
