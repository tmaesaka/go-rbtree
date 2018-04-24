package rbtree

import (
	"sort"
	"strconv"
	"strings"
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

func TestUpdate(t *testing.T) {
	tree := NewTree()
	testKey := []byte("apricot")
	testVal := "yummy cobbler"

	if err := tree.Update([]byte("apricot"), "new value"); err == nil {
		t.Errorf("expected node to not exist")
	}

	tree.Insert(testKey, "jam")
	tree.Insert([]byte("banana"), "smoothie")
	tree.Insert([]byte("clementine"), "cake")

	if err := tree.Update(testKey, testVal); err != nil {
		t.Error(err)
	}

	n, _ := tree.Find(testKey)
	if n.value != testVal {
		t.Errorf("expected: %s, got: %s", testVal, n.value)
	}
}

func TestInorder(t *testing.T) {
	tree := NewTree()
	keys := []string{"x", "c", "a", "y", "b", "z"}

	for _, key := range keys {
		tree.Insert([]byte(key), nil)
	}

	i := 0
	sort.Strings(keys)
	tree.Inorder(func(k []byte, v interface{}) {
		if strings.Compare(keys[i], string(k)) != 0 {
			t.Errorf("expected: %s, got: %s", keys[i], string(k))
		}
		i++
	})
}
