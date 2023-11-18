// Package bst is a Golang implementation of a binary search tree, utilizing
// as much recursion as possible for code readability and compactness
package bst

import "fmt"

// Base node structure for the tree
type Node struct {
	key         int
	val         string
	left, right *Node
}

// Insert recursively adds a node to the BST by traversing it, finding the
// first free location
func (n *Node) insert(key int, val string) *Node {
	// If the current node does not exist, it's created and returned
	// to the calling method, which adds it to the tree
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

// Traversal is used to print the tree, with the corresponding level of the
// records found
func (n *Node) traversal(lvl int) {
	if n != nil {
		n.left.traversal(lvl + 1)
		fmt.Println(lvl, n.key, n.val)
		n.right.traversal(lvl + 1)
	}
}
