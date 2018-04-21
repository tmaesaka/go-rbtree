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

// Height returns the height (maximum depth) of the Tree.
func (tree *Tree) Height() uint {
	tree.mtx.RLock()
	defer tree.mtx.RUnlock()

	var height uint
	if tree.root == nil {
		return height
	}

	// TODO(toru): Try using buffered channel as a queue.
	queue := []*node{tree.root}
	var curr *node

	for len(queue) > 0 {
		size := len(queue)
		for i := size; i > 0; i-- {
			// Shift
			curr, queue = queue[0], queue[1:]
			if curr.left != nil {
				queue = append(queue, curr.left)
			}
			if curr.right != nil {
				queue = append(queue, curr.right)
			}
		}
		height++
	}
	return height
}

// Find returns a node that matches the given key, otherwise nil.
// FIXME(toru): This function should return an iterator.
func (tree *Tree) Find(key []byte) (*node, bool) {
	tree.mtx.RLock()
	n, ok := tree.find(key)
	tree.mtx.RUnlock()

	return n, ok
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

	// Next step is to fix any Red-Black Tree property violations caused
	// by the node insertion. Work our way up the tree from the node.
	curr := n
	for {
		// Case 1: Root node. Paint it black and we're done.
		if curr.parent == nil {
			curr.color = Black
			break
		}

		// Case 2: Parent node is painted black, nothing to do.
		if curr.parent.color == Black {
			break
		}

		// Pointers to grasp the state of the current subtree.
		uncle := curr.uncle()
		grandparent := curr.grandparent()

		// Case 3: Parent and Uncle nodes are painted red.
		if uncle != nil && uncle.color == Red {
			curr.parent.color = Black
			uncle.color = Black
			grandparent.color = Red

			// Work our way up the tree.
			curr = grandparent
			continue
		}

		// Case 4: Getting here means the parent is painted red and the uncle is
		// painted black. Rotate the branch if the current node is an inner-node.
		if grandparent.left != nil && curr == grandparent.left.right {
			tree.leftRotate(curr.parent)
			curr = curr.left
		} else if grandparent.right != nil && curr == grandparent.right.left {
			tree.rightRotate(curr.parent)
			curr = curr.right
		}

		// Getting here means it's safe to run rotation in order to make the previous
		// parent become the parent of both curr and former grandparent.
		if curr == curr.parent.left {
			tree.rightRotate(grandparent)
		} else {
			tree.leftRotate(grandparent)
		}
		curr.parent.color = Black
		grandparent.color = Red
	}

	return nil
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

func (tree *Tree) replaceNode(oldNode, newNode *node) {
	if oldNode.parent == nil {
		tree.root = newNode
	} else if oldNode == oldNode.parent.left {
		oldNode.parent.left = newNode
	} else {
		oldNode.parent.right = newNode
	}

	if newNode != nil {
		newNode.parent = oldNode.parent
	}
}

func (tree *Tree) leftRotate(n *node) {
	y := n.right
	n.right = y.left

	if y.left != nil {
		y.left.parent = n
	}
	y.parent = n.parent

	if n.parent == nil {
		tree.root = y
	} else if n == n.parent.left {
		n.parent.left = y
	} else {
		n.parent.right = y
	}
	y.left = n
	n.parent = y
}

func (tree *Tree) rightRotate(n *node) {
	y := n.left
	n.left = y.right

	if y.right != nil {
		y.right.parent = n
	}
	y.parent = n.parent

	if n.parent == nil {
		tree.root = y
	} else if n == n.parent.left {
		n.parent.left = y
	} else {
		n.parent.right = y
	}
	y.right = n
	n.parent = y
}

func (tree *Tree) find(key []byte) (*node, bool) {
	var found bool
	curr := tree.root
	for curr != nil {
		cmp := bytes.Compare(key, curr.key)
		if cmp < 0 {
			curr = curr.left
		} else if cmp > 0 {
			curr = curr.right
		} else {
			found = true
			break
		}
	}
	return curr, found
}
