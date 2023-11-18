package bst

type Node struct {
	key         int
	val         string
	left, right *Node
}

func (n *Node) insert(key int, val string) *Node {
	if n == nil {
		new_node := Node{key, val, nil, nil}
		return &new_node
	}

	if n.key > key {
		n.left = n.left.insert(key, val)
	} else if n.key < key {
		n.right = n.right.insert(key, val)
	}

	return n
}
