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
