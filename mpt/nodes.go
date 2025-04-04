package mpt

type Node interface {
	Bytes() []byte
}

func IsEmptyNode(node Node) bool {
	return node == nil
}
