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
func (n *Node) Insert(key int, val string) *Node {
	// If the current node does not exist, it's created and returned
	// to the calling method, which adds it to the tree
	if n == nil {
		new_node := Node{key, val, nil, nil}
		return &new_node
	}

	if n.key > key {
		n.left = n.left.Insert(key, val)
	} else if n.key < key {
		n.right = n.right.Insert(key, val)
	}

	return n
}

// Search searches the tree for a key and returns the node associated with that
// key and the depth of the node. If the node does not exists, nil is returned as the node
func (n *Node) Search(key int) (*Node, int) {
	depth := 0
	for {
		if n == nil {
			return nil, depth
		} else if n.key > key {
			n = n.left
		} else if n.key < key {
			n = n.right
		} else if n.key == key {
			return n, depth
		}
		depth += 1
	}
}

// findMin finds the minimal value from a starting node
func (n *Node) FindMin() *Node {
	for {
		if n.left != nil {
			n = n.left
		} else {
			return n
		}
	}
}

// delete removes nodes from the tree and moves nodes around as needed
func (n *Node) Delete(key int) *Node {
	if n == nil {
		return nil
	} else if n.key > key {
		n.left = n.left.Delete(key)
	} else if n.key < key {
		n.right = n.right.Delete(key)
	} else {
		// No children
		if n.left == nil && n.right == nil {
			n = nil

			// One child
		} else if n.left == nil {
			n = n.right
		} else if n.right == nil {
			n = n.left

			// Two children
		} else {
			node_min_right := n.right.FindMin()
			n.key = node_min_right.key
			n.val = node_min_right.val

			n.right.Delete(n.key)
		}
	}

	return n
}

// Traversal is used to print the tree, with the corresponding level of the
// records found
func (n *Node) Traversal(lvl int) {
	if n != nil {
		n.left.Traversal(lvl + 1)
		fmt.Println(lvl, n.key, n.val)
		n.right.Traversal(lvl + 1)
	}
}
