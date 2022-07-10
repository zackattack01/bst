package bst

type (
	Node struct {
		// if present, the left node must contain a value less than that of the current node
		left *Node
		// if present, the right node must contain a value greater than that of the current node
		right *Node
		// all nodes beyond the root will have a parent node, indicating the last node used for comparison
		parent *Node
		value  int
	}
)

func (n *Node) IntValue() int {
	return n.value
}

func IntegerNodeValues(nodes []*Node) []int {
	outputNodes := make([]int, 0)
	for _, node := range nodes {
		outputNodes = append(outputNodes, node.value)
	}

	return outputNodes
}

// findReplacement is called on a Node pending deletion
// returns the appropriate Node to take its place
func (n *Node) findReplacement() *Node {
	if n == nil {
		return nil
	}

	// if there is no left, it's safe to go right, even if nil (valid case for a leaf Node)
	if n.left == nil {
		return n.right
	}

	// if there is no right, we know we have a left at this point so use that
	if n.right == nil {
		return n.left
	}

	// we are a Node with both a left and a right child, find the smallest value greater than
	// our own to use as the successor
	return n.right.min()
}

// max returns the Node with the maximum value from this Node's point in the tree
func (n *Node) max() *Node {
	currentNode := n
	for currentNode.right != nil {
		currentNode = currentNode.right
	}

	return currentNode
}

// min returns the Node with the minimum value from this Node's point in the tree
func (n *Node) min() *Node {
	currentNode := n
	for currentNode.left != nil {
		currentNode = currentNode.left
	}

	return currentNode
}

// handful of private readability/convenience methods to remove boilerplate elsewhere ---------
// isRoot - only the root Node has a nil parent
func (n *Node) isRoot() bool      { return n.parent == nil }
func (n *Node) isLeftChild() bool { return !n.isRoot() && n.parent.left == n }

// isFork - does this Node have both a left and right child
func (n *Node) isFork() bool { return n.left != nil && n.right != nil }
