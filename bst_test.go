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
