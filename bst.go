// Usage:
//
//     tree := NewIntBST([]int{26, 82, 16, 92, 33})
//     tree.Delete(tree.SearchInt(16))
//     nodes, maxDepth := tree.DeepestNodes())
//     fmt.Printf("%v found at depth %d", IntegerNodeValues(nodes), maxDepth)
//
// See bin/standards for how to regenerate this file and run tests
package bst

type (
	// BST provides the exported interface wrapping all of our internal tree logic.
	BST interface {
		// DeepestNodes returns the Node values for the corresponding deepest nodes
		DeepestNodes() ([]*Node, int)

		// Delete ensures removal of the given IntValue,
		// returning true if it was present to begin with or false if no action was required
		Delete(node *Node) bool
		// Insert appends the given IntValue as a leaf Node after traversing the tree to find the appropriate location.
		// returns true if the IntValue is new and insertion was required, false if no action was required
		Insert(node *Node) bool
		Max() *Node
		Min() *Node
		// SearchInt attempts to retrieve a Node containing the give integer searchValue
		// returns nil if the IntValue does not exist
		SearchInt(searchValue int) *Node
		WalkInOrder() []*Node   // left subtree, root, and then right subtree
		WalkPostOrder() []*Node // left subtree, then right subtree, then root
		WalkPreOrder() []*Node  // root, left subtree, then right subtree
	}

	visitedNode struct {
		depth int
		node  *Node
	}

	tree struct {
		root *Node
	}
)

// NewIntBST returns a tree object adhering to the BST interface
func NewIntBST(values []int) BST {
	t := tree{}
	for _, value := range values {
		t.Insert(&Node{value: value})
	}

	return &t
}

func (t *tree) Delete(removal *Node) bool {
	if removal == nil {
		return false
	}

	replacement := removal.findReplacement()
	// breaking out this case because it is more complicated than the others-
	// if the Node to be removed has both a left and right child, the replacement
	// will be the lowest IntValue exceeding its own, and not necessarily a direct parent (see findReplacement)
	if removal.isFork() {
		if replacement.parent != removal {
			t.shiftNodes(replacement, replacement.right)
			replacement.right = removal.right
			replacement.right.parent = replacement
		}

		t.shiftNodes(removal, replacement)
		// we must additionally replace the left subtree of the deleted Node here to complete
		// replacement of a fork Node
		replacement.left = removal.left
		replacement.left.parent = replacement
	} else {
		// any leaf or single child nodes are handled by simply shifting the next available
		// IntValue as right || left || nil
		t.shiftNodes(removal, replacement)
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
			return false // always safe to return (no action needed) if the IntValue already exists
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

func (t *tree) Min() *Node { return t.root.min() }
func (t *tree) Max() *Node { return t.root.max() }

func (t *tree) SearchInt(searchValue int) *Node {
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

// shiftNodes is a utility method for use by the Node Delete process
// this will delete the Node for removal and shift the replacement into its place
func (t *tree) shiftNodes(removal *Node, replacement *Node) {
	if removal == nil {
		return
	}

	if removal.isRoot() {
		t.root = replacement
		return
	}

	// update parent pointers according to our removal Node's current relationship
	if removal.isLeftChild() {
		removal.parent.left = replacement
	} else {
		removal.parent.right = replacement
	}

	// ensure the replacements parent pointer is updated if required
	if replacement != nil {
		replacement.parent = removal.parent
	}
}
