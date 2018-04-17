package rbtree

import (
	"strconv"
	"testing"
)

func TestLen(t *testing.T) {
	expected := 256
	tree := NewTree()

	for i := 0; i < expected; i++ {
		tree.Insert([]byte(strconv.Itoa(i)), "value")
	}

	if uint(expected) != tree.Len() {
		t.Errorf("Expected: %d, Got: %d", expected, tree.Len())
	}
}
