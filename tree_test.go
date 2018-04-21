package rbtree

import (
	"strconv"
	"testing"
)

var nodeColorLabel = [2]string{"red", "black"}

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

func TestHeight(t *testing.T) {
	tree := NewTree()

	height := tree.Height()
	if height != 0 {
		t.Errorf("expected: 0, got: %d", height)
	}

	for i := 97; i <= 122; i++ {
		tree.Insert([]byte(string(i)), nil)
	}

	height = tree.Height()
	if height != 7 {
		t.Errorf("expected: 7, got: %d", height)
	}
}

func TestFind(t *testing.T) {
	tree := NewTree()
	tree.Insert([]byte("apple"), "sauce")

	node, _ := tree.Find([]byte("apple"))
	if node.value != "sauce" {
		t.Errorf("expected sauce, got: %s", node.value)
	}

	node, ok := tree.Find([]byte("banana"))
	if ok {
		t.Errorf("expected node to not exist, got: %v", node)
	}
}

func TestInsert(t *testing.T) {
	t.Run("sequence input", func(t *testing.T) {
		tree := NewTree()
		tree.Insert([]byte("aaa"), nil)
		tree.Insert([]byte("bbb"), nil)
		tree.Insert([]byte("ccc"), nil)
		tree.Insert([]byte("ddd"), nil)

		l := tree.root.left
		r := tree.root.right
		rr := tree.root.right.right

		if string(tree.root.key) != "bbb" || tree.root.color != Black {
			t.Errorf("expected bbb (black), got: %s (%s)",
				tree.root.key, nodeColorLabel[tree.root.color])
		}
		if string(l.key) != "aaa" || l.color != Black {
			t.Errorf("expected aaa (black), got: %s (%s)",
				l.key, nodeColorLabel[l.color])
		}
		if string(r.key) != "ccc" || r.color != Black {
			t.Errorf("expected ccc (black), got: %s (%s)",
				r.key, nodeColorLabel[r.color])
		}
		if string(rr.key) != "ddd" || rr.color != Red {
			t.Errorf("expected ddd (red), got: %s (%s)",
				rr.key, nodeColorLabel[rr.color])
		}
	})
}
