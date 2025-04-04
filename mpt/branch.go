package mpt

type BranchNode struct {
	Branches [16]Node
	value    any
}

var _ Node = (*BranchNode)(nil)

func NewBranchNode() *BranchNode {
	return &BranchNode{
		Branches: [16]Node{},
	}
}

func (b *BranchNode) Value() any {
	return b.value
}

func (b *BranchNode) SetBranch(nibble Nibble, node Node) {
	b.Branches[int(nibble)] = node
}

func (b *BranchNode) RemoveBranch(nibble Nibble) {
	b.Branches[int(nibble)] = nil
}

func (b *BranchNode) SetValue(value any) {
	b.value = value
}

func (b *BranchNode) RemoveValue() {
	b.value = nil
}

func (b BranchNode) HasValue() bool {
	return b.value != nil
}
