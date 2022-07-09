package bst

import "fmt"

type (
	Node struct {
		// if present, the left Node must contain a value less than that of the current Node
		left *Node
		// if present, the right Node must contain a value greater than that of the current Node
		right *Node
		// all nodes beyond the root will have a parent Node, indicating the last Node used for comparison
		parent *Node
		value  int
	}
)

func NewIntegerNode(value int) *Node {
	return &Node{value: value}
}

func NodesToIntSlice(nodes []*Node) []int {
	outputNodes := make([]int, 0)
	for _, node := range nodes {
		outputNodes = append(outputNodes, node.value)
	}

	return outputNodes
}

// successor returns the Node with the smallest value that is still greater than the called Node's value
func (n *Node) successor() *Node {
	if n == nil {
		return nil
	} else if n.right != nil {
		return n.right.min()
	}

	currentNode := n
	parentNode := n.parent
	for parentNode != nil && currentNode == parentNode.right {
		currentNode = parentNode
		parentNode = parentNode.parent
	}

	return parentNode
}

// max returns the Node with the maximum value from this Node's point in the tree
func (n *Node) max() *Node {
	currentNode := n
	for {
		if currentNode.right == nil {
			return currentNode
		}

		currentNode = currentNode.right
	}
}

// min returns the Node with the minimum value from this Node's point in the tree
func (n *Node) min() *Node {
	currentNode := n
	for {
		if currentNode.left == nil {
			return currentNode
		}

		currentNode = currentNode.left
	}
}

// readability/convenience method - only the root Node has a nil parent
func (n *Node) isRoot() bool   { return n.parent == nil }
func (n *Node) String() string { return fmt.Sprintf("%v", n.value) }
