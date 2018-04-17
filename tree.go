// Package rbtree implements the self balancing Red-Black Tree
// data structure with a concise key-value interface.
package rbtree

import (
	"bytes"
	"fmt"
	"sync"
)

// Tree represents a Red-Black Tree. Use the public API of this structure
// to perform operations against the Tree.
type Tree struct {
	root  *node
	mtx   sync.RWMutex
	count uint
}

// NewTree creates and returns a new Red-Black Tree.
func NewTree() *Tree {
	t := Tree{}
	return &t
}

// Len returns the number of nodes currently in the Tree.
func (tree *Tree) Len() uint {
	tree.mtx.RLock()
	count := tree.count
	tree.mtx.RUnlock()

	return count
}

// Insert adds a new node to the Tree, indexed by the given key.
func (tree *Tree) Insert(key []byte, value interface{}) error {
	tree.mtx.Lock()
	defer tree.mtx.Unlock()

	n := &node{key: key, value: value}

	// First step is to do a standard BST insertion.
	if err := tree.bstInsert(n); err != nil {
		return err
	}

	return fmt.Errorf("Under Construction")
}

// Delete removes a node from the Tree that matches the given key.
func (tree *Tree) Delete(key []byte) error {
	return fmt.Errorf("Unimplemented")
}

func (tree *Tree) bstInsert(n *node) error {
	curr := tree.root

	// Tree is empty, tweak the node as root.
	if curr == nil {
		n.color = Black
		tree.root = n
		tree.count++
		return nil
	}

	var parent *node

	// Iteratively search for a parent to dangle from.
	for curr != nil {
		parent = curr

		// TODO(toru): Support user-supplied comparator function.
		cmp := bytes.Compare(n.key, curr.key)

		if cmp < 0 {
			curr = curr.left
		} else if cmp > 0 {
			curr = curr.right
		} else {
			// TODO(toru): Support duplicate keys like many tree-based key-value
			// databases do. Until then, duplication is an error for simplicity.
			return fmt.Errorf("duplicate key")
		}
	}

	// Found the parent. Link it to the node and vice versa.
	n.parent = parent
	if bytes.Compare(n.key, parent.key) < 0 {
		parent.left = n
	} else {
		parent.right = n
	}

	tree.count++
	return nil
}

func (tree *Tree) balance(n *node) {
}

func (tree *Tree) leftRotate(n *node) {
}

func (tree *Tree) rightRotate(n *node) {
}
