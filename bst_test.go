package bst

import (
	"testing"
)

func TestInsert(t *testing.T) {
	got := insert()
	if !got {
		t.Error("Insert returned false")
	}
}

func TestNodeCreation(t *testing.T) {
	newNode := Node{key: 1, val: "data", left: nil, right: nil}

	if newNode.key != 1 || newNode.val != "data" {
		t.Errorf("Expected key: 1, val: data. Received key: %v, val: %v", newNode.key, newNode.val)
	}
}
