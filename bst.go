package bst

type (
	// BST provides the exported interface wrapping all of our internal data/state.
	// we will define all exported (public) methods on this type to maximize
	// operational flexibility going forward
	BST interface {
		// Insert appends the given value as a leaf Node after traversing the tree to find the appropriate location.
		// returns true if the value is new and insertion was required, false if no action was required
		Insert(node *Node) bool
		// Delete ensures removal of the given value,
		//returning true if it was present to begin with or false if no action was required
		Delete(node *Node) bool
		Min() *Node
		Max() *Node
		// DeepestNodes returns the integer values for the corresponding deepest nodes
		DeepestNodes() ([]*Node, int)
		Find(searchValue int) *Node
	}

	tree struct {
		root *Node
	}
)

func NewFromIntSlice(values []int) BST {
	t := tree{}
	for _, value := range values {
		t.Insert(NewIntegerNode(value))
	}

	return &t
}

func (t *tree) DeepestNodes() ([]*Node, int) {
	deepestNodes := make([]*Node, 0)
	maxDepth := 0
	currentDepth := 0
	currentNode := t.root
	type visitedNode struct {
		depth int
		node  *Node
	}
	visitedNodes := make([]*visitedNode, 0)
	for {
		if currentNode != nil {
			visitedNodes = append(visitedNodes, &visitedNode{currentDepth, currentNode})
			currentNode = currentNode.left
		} else if len(visitedNodes) != 0 {
			lastVisitedIndex := len(visitedNodes) - 1
			latestVisited := visitedNodes[lastVisitedIndex]
			visitedNodes = visitedNodes[:lastVisitedIndex]
			currentNode = latestVisited.node.right
			currentDepth = latestVisited.depth // reset the depth counter as required for this level
		} else {
			// current Node is nil and there is nothing left to go back to, we are done
			return deepestNodes, maxDepth
		}

		if currentNode == nil {
			continue // prevent premature depth updates in case we've just navigated to a nil level
		}

		currentDepth++

		if currentDepth == maxDepth {
			deepestNodes = append(deepestNodes, currentNode)
		} else if currentDepth > maxDepth {
			deepestNodes = []*Node{currentNode}
			maxDepth = currentDepth
		}
	}
}

func (t *tree) Min() *Node { return t.root.min() }
func (t *tree) Max() *Node { return t.root.max() }

func (t *tree) Delete(targetDeletion *Node) bool {
	if targetDeletion.left == nil {
		t.shiftNodes(targetDeletion, targetDeletion.right)
	} else if targetDeletion.right == nil {
		t.shiftNodes(targetDeletion, targetDeletion.left)
	} else {
		successor := targetDeletion.successor()
		if successor.parent != targetDeletion {
			t.shiftNodes(successor, successor.right)
			successor.right = targetDeletion.right
			successor.right.parent = successor
		}

		t.shiftNodes(targetDeletion, successor)
		successor.left = targetDeletion.left
		successor.left.parent = successor
	}

	return true
}

func (t *tree) Insert(newNode *Node) bool {
	var parentNode *Node // trailing pointer for iterative search
	currentNode := t.root

	// handle case for uninitialized root and return immediately, no need to complicate the remaining logic
	if t.root == nil {
		t.root = newNode
		return true
	}

	// traverse the tree, maintaining a trailing pointer to the parent Node
	for currentNode != nil {
		parentNode = currentNode
		if newNode.value == currentNode.value {
			return false // always safe to return (no action needed) if the value already exists
		}

		if newNode.value < currentNode.value {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}

	// at this point we know that we've reached the furthest relevant branch,
	// do one final comparison then it's time to leaf
	newNode.parent = parentNode
	if newNode.value < parentNode.value {
		parentNode.left = newNode
	} else {
		parentNode.right = newNode
	}

	return true
}

func (t *tree) Find(searchValue int) *Node {
	currentNode := t.root

	for currentNode != nil && searchValue != currentNode.value {
		if searchValue < currentNode.value {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}

	return currentNode
}

func (t *tree) shiftNodes(oldNode *Node, newNode *Node) {
	if oldNode.isRoot() {
		t.root = newNode
		return
	}

	if oldNode == oldNode.parent.left {
		oldNode.parent.left = newNode
	} else {
		oldNode.parent.right = newNode
	}

	if newNode != nil {
		newNode.parent = oldNode.parent
	}
}
