package bst

import (
	"testing"
)

func TestInsert(t *testing.T) {
	var root *Node
	root = root.Insert(5, "data")
	root = root.Insert(10, "data")
	root = root.Insert(1, "data")

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
	root = root.Insert(5, "data")
	root = root.Insert(10, "data1")
	root = root.Insert(1, "data2")

	should_fail, _ := root.Search(0)
	if should_fail != nil {
		t.Errorf("Expected error to be raise, received: %v", should_fail)
	}

	should_pass_root, _ := root.Search(5)
	if should_pass_root.val != "data" {
		t.Errorf("Search root: Expected \"data\", received: %v", should_pass_root)
	}

	should_pass_right, _ := root.Search(10)
	if should_pass_right.val != "data1" {
		t.Errorf("Search right: Expected \"data\", received: %v", should_pass_right)
	}

	should_pass_left, _ := root.Search(1)
	if should_pass_left.val != "data2" {
		t.Errorf("Search left: Expected \"data\", received: %v", should_pass_left)
	}
}

func TestTraversal(t *testing.T) {
	var root *Node
	root = root.Insert(5, "data")
	root = root.Insert(10, "data")
	root = root.Insert(1, "data")

	root.Traversal(0)
}

func TestDeleteNoChildren(t *testing.T) {
	var root *Node
	root = root.Insert(60, "data")
	root = root.Insert(55, "data")
	root = root.Insert(68, "data")

	root = root.Delete(55)

	got, _ := root.Search(55)
	if got != nil {
		t.Errorf("Expected: nil, Got: %v", got)
	}
}

func TestDeleteOneChild(t *testing.T) {
	var root *Node
	root = root.Insert(60, "data")
	root = root.Insert(55, "data")
	root = root.Insert(68, "data")
	root = root.Insert(66, "data")
	root = root.Insert(69, "data")
	root = root.Insert(62, "data")
	root = root.Insert(67, "data")
	root = root.Insert(63, "data")

	root = root.Delete(62)

	got, _ := root.Search(62)
	if got != nil {
		t.Errorf("Expected: nil, Got: %v", got)
	}

	got, depth := root.Search(63)
	if got == nil || depth != 3 {
		t.Errorf("Expected: struct, 3, Got: %v, %v", got, depth)
	}
}

func TestDeleteTwoChildren(t *testing.T) {
	var root *Node
	root = root.Insert(60, "data")
	root = root.Insert(55, "data")
	root = root.Insert(68, "data")
	root = root.Insert(66, "data")
	root = root.Insert(69, "data")
	root = root.Insert(62, "data")
	root = root.Insert(67, "data")
	root = root.Insert(63, "data")

	root = root.Delete(60)

	got, _ := root.Search(60)
	if got != nil {
		t.Errorf("Expected: nil, Got: %v", got)
	}

	got, depth := root.Search(62)
	if got == nil || depth != 0 {
		t.Errorf("Expected: nil, 0, Got: %v, %v", got, depth)
	}

	got, depth = root.Search(63)
	if got == nil || depth != 3 {
		t.Errorf("Expected: nil, 3, Got: %v, %v", got, depth)
	}
}

func TestFindMin(t *testing.T) {
	var root *Node
	root = root.Insert(60, "data")
	root = root.Insert(55, "data")
	root = root.Insert(68, "data")
	root = root.Insert(66, "data")
	root = root.Insert(69, "data")
	root = root.Insert(62, "data")
	root = root.Insert(67, "data")
	root = root.Insert(63, "data")

	root_node_min := root.FindMin().key
	if root_node_min != 55 {
		t.Errorf("Expected: 55, Got: %v", root_node_min)
	}

	non_root_node, _ := root.Search(68)
	non_root_min := non_root_node.FindMin().key
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
