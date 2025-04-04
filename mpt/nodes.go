package mpt

type Node interface {
	Value() any
}

func IsEmptyNode(node Node) bool {
	return node == nil
}
