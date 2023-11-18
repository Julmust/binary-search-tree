package bst

import (
	"testing"
)

func TestInsert(t *testing.T) {
	var root *Node
	root = root.insert(5, "data")
	root = root.insert(10, "data")
	root = root.insert(1, "data")

	if root.key != 5 || root.val != "data" {
		t.Error("Insert failed")
	}

	if root.right.key != 10 {
		t.Error("Right insert failed")
	}

	if root.left.key != 1 {
		t.Error("Left insert failed")
	}
}

func TestSearch(t *testing.T) {
	var root *Node
	root = root.insert(5, "data")
	root = root.insert(10, "data")
	root = root.insert(1, "data")

	should_fail := root.search(0)
	if should_fail != "Key not found" {
		t.Errorf("Expected error to be raise, received: %v", should_fail)
	}

	should_pass_root := root.search(5)
	if should_pass_root != "data" {
		t.Errorf("Search root: Expected \"data\", received: %v", should_pass_root)
	}

	should_pass_left := root.search(1)
	if should_pass_left != "data" {
		t.Errorf("Search left: Expected \"data\", received: %v", should_pass_left)
	}

	should_pass_right := root.search(10)
	if should_pass_right != "data" {
		t.Errorf("Search right: Expected \"data\", received: %v", should_pass_right)
	}
}

func TestTraversal(t *testing.T) {
	var root *Node
	root = root.insert(5, "data")
	root = root.insert(10, "data")
	root = root.insert(1, "data")

	root.traversal(0)
}

func TestNodeCreation(t *testing.T) {
	newNode := Node{key: 1, val: "data", left: nil, right: nil}

	if newNode.key != 1 || newNode.val != "data" {
		t.Errorf("Expected key: 1, val: data. Received key: %v, val: %v", newNode.key, newNode.val)
	}
}
