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

func TestNodeCreation(t *testing.T) {
	newNode := Node{key: 1, val: "data", left: nil, right: nil}

	if newNode.key != 1 || newNode.val != "data" {
		t.Errorf("Expected key: 1, val: data. Received key: %v, val: %v", newNode.key, newNode.val)
	}
}
