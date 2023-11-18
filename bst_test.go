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
	root = root.insert(10, "data1")
	root = root.insert(1, "data2")

	should_fail, _ := root.search(0)
	if should_fail != nil {
		t.Errorf("Expected error to be raise, received: %v", should_fail)
	}

	should_pass_root, _ := root.search(5)
	if should_pass_root.val != "data" {
		t.Errorf("Search root: Expected \"data\", received: %v", should_pass_root)
	}

	should_pass_right, _ := root.search(10)
	if should_pass_right.val != "data1" {
		t.Errorf("Search right: Expected \"data\", received: %v", should_pass_right)
	}

	should_pass_left, _ := root.search(1)
	if should_pass_left.val != "data2" {
		t.Errorf("Search left: Expected \"data\", received: %v", should_pass_left)
	}
}

func TestTraversal(t *testing.T) {
	var root *Node
	root = root.insert(5, "data")
	root = root.insert(10, "data")
	root = root.insert(1, "data")

	root.traversal(0)
}

func TestDeleteNoChildren(t *testing.T) {
	var root *Node
	root = root.insert(60, "data")
	root = root.insert(55, "data")
	root = root.insert(68, "data")

	root = root.delete(55)

	got, _ := root.search(55)
	if got != nil {
		t.Errorf("Expected: nil, Got: %v", got)
	}
}

func TestDeleteOneChild(t *testing.T) {
	var root *Node
	root = root.insert(60, "data")
	root = root.insert(55, "data")
	root = root.insert(68, "data")
	root = root.insert(66, "data")
	root = root.insert(69, "data")
	root = root.insert(62, "data")
	root = root.insert(67, "data")
	root = root.insert(63, "data")

	root = root.delete(62)

	got, _ := root.search(62)
	if got != nil {
		t.Errorf("Expected: nil, Got: %v", got)
	}

	got, depth := root.search(63)
	if got == nil || depth != 3 {
		t.Errorf("Expected: struct, 3, Got: %v, %v", got, depth)
	}
}

func TestDeleteTwoChildren(t *testing.T) {
	var root *Node
	root = root.insert(60, "data")
	root = root.insert(55, "data")
	root = root.insert(68, "data")
	root = root.insert(66, "data")
	root = root.insert(69, "data")
	root = root.insert(62, "data")
	root = root.insert(67, "data")
	root = root.insert(63, "data")

	root = root.delete(60)

	got, _ := root.search(60)
	if got != nil {
		t.Errorf("Expected: nil, Got: %v", got)
	}

	got, depth := root.search(62)
	if got == nil || depth != 0 {
		t.Errorf("Expected: nil, 0, Got: %v, %v", got, depth)
	}

	got, depth = root.search(63)
	if got == nil || depth != 3 {
		t.Errorf("Expected: nil, 3, Got: %v, %v", got, depth)
	}
}

func TestFindMin(t *testing.T) {
	var root *Node
	root = root.insert(60, "data")
	root = root.insert(55, "data")
	root = root.insert(68, "data")
	root = root.insert(66, "data")
	root = root.insert(69, "data")
	root = root.insert(62, "data")
	root = root.insert(67, "data")
	root = root.insert(63, "data")

	root_node_min := root.findMin().key
	if root_node_min != 55 {
		t.Errorf("Expected: 55, Got: %v", root_node_min)
	}

	non_root_node, _ := root.search(68)
	non_root_min := non_root_node.findMin().key
	if non_root_min != 62 {
		t.Errorf("Expected: 62, Got: %v", non_root_min)
	}
}

func TestNodeCreation(t *testing.T) {
	newNode := Node{key: 1, val: "data", left: nil, right: nil}

	if newNode.key != 1 || newNode.val != "data" {
		t.Errorf("Expected key: 1, val: data. Received key: %v, val: %v", newNode.key, newNode.val)
	}
}
