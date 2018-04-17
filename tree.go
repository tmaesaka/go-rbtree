// Package rbtree implements the self balancing Red-Black Tree
// data structure with a concise key-value interface.
package rbtree

import (
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
	return fmt.Errorf("Unimplemented")
}

// Delete removes a node from the Tree that matches the given key.
func (tree *Tree) Delete(key []byte) error {
	return fmt.Errorf("Unimplemented")
}

func (tree *Tree) balance(n *node) {
}

func (tree *Tree) leftRotate(n *node) {
}

func (tree *Tree) rightRotate(n *node) {
}
