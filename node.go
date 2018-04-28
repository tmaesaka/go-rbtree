package rbtree

type Color uint8

const (
	Red = Color(iota)
	Black
)

// node is the building block of a tree and holds various data together.
type node struct {
	color       Color
	parent      *node
	left, right *node
	key         []byte
	value       interface{}
}

// grandparent returns a pointer to a grandparent node, or nil if there is none.
func (n *node) grandparent() *node {
	if n.parent == nil {
		return nil
	}
	return n.parent.parent
}

// siblings returns a pointer to a sibling node, or nil if there is none.
func (n *node) sibling() *node {
	if n.parent == nil {
		return nil
	}

	if n == n.parent.left {
		return n.parent.right
	} else {
		return n.parent.left
	}
}

// uncle returns a pointer to an uncle node, or nil if there is none.
func (n *node) uncle() *node {
	g := n.grandparent()
	if g == nil {
		return nil
	}
	return n.parent.sibling()
}

// isLeaf returns whether the node is a leaf node.
func (n *node) isLeaf() bool {
	return n.left == nil && n.right == nil
}

// minNode returns a pointer to the minimum node in the subtree.
func (n *node) minNode() *node {
	curr := n
	for curr.left != nil {
		curr = curr.left
	}
	return curr
}
