package mpt

type BranchNode struct {
	Branches [16]Node
	Value    []byte
}

var _ Node = (*BranchNode)(nil)

func NewBranchNode() *BranchNode {
	return &BranchNode{
		Branches: [16]Node{},
	}
}

func (b *BranchNode) Bytes() []byte {
	return b.Value
}

func (b *BranchNode) SetBranch(nibble Nibble, node Node) {
	b.Branches[int(nibble)] = node
}

func (b *BranchNode) RemoveBranch(nibble Nibble) {
	b.Branches[int(nibble)] = nil
}

func (b *BranchNode) SetValue(value []byte) {
	b.Value = value
}

func (b *BranchNode) RemoveValue() {
	b.Value = nil
}

func (b BranchNode) HasValue() bool {
	return b.Value != nil
}
