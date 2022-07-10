package bst

func (t *tree) DeepestNodes() ([]*Node, int) {
	deepestNodes := make([]*Node, 0)

	if t.root == nil {
		return deepestNodes, 0
	}
	// we'll do in order (shouldn't matter). some duplication here
	// to avoid complicating the exposed Walk methods as there doesn't seem
	// to be sufficient use case for introducing callbacks here yet
	rootVisit := &visitedNode{0, t.root}
	visited := t.walkFromNode(t.root.left, 1)
	visited = append(visited, rootVisit)
	visited = append(visited, t.walkFromNode(t.root.right, 1)...)
	maxDepthSeen := 0
	for _, vnode := range visited {
		if vnode.depth == maxDepthSeen {
			deepestNodes = append(deepestNodes, vnode.node)
		} else if vnode.depth > maxDepthSeen {
			maxDepthSeen = vnode.depth
			deepestNodes = []*Node{vnode.node}
		}
	}

	return deepestNodes, maxDepthSeen
}

func (t *tree) WalkInOrder() []*Node {
	rootVisit := &visitedNode{0, t.root}
	visited := t.walkFromNode(t.root.left, 1)
	visited = append(visited, rootVisit)
	visited = append(visited, t.walkFromNode(t.root.right, 1)...)
	return extractNodesFromVisited(visited)
}

func (t *tree) WalkPostOrder() []*Node {
	rootVisit := &visitedNode{0, t.root}
	visited := t.walkFromNode(t.root.left, 1)
	visited = append(visited, t.walkFromNode(t.root.right, 1)...)
	visited = append(visited, rootVisit)
	return extractNodesFromVisited(visited)
}

func (t *tree) WalkPreOrder() []*Node {
	rootVisit := &visitedNode{0, t.root}
	visited := []*visitedNode{rootVisit}
	visited = append(visited, t.walkFromNode(t.root.left, 1)...)
	visited = append(visited, t.walkFromNode(t.root.right, 1)...)
	return extractNodesFromVisited(visited)
}

func (t *tree) walkFromNode(current *Node, depth int) []*visitedNode {
	visited := make([]*visitedNode, 0)
	queue := make([]*visitedNode, 0)

	// continue until there is nothing being visited or pending in queue
	for current != nil || len(queue) > 0 {
		if current != nil { // if present, add ourselves to the queue and move left
			queue = append(queue, &visitedNode{depth, current})
			visited = append(visited, &visitedNode{depth, current})
			current = current.left
		} else if len(queue) > 0 { // otherwise left wasn't there, go back to the queue and try right
			lastQueuedIndex := len(queue) - 1
			latestVisit := queue[lastQueuedIndex]
			queue = queue[:lastQueuedIndex]
			current = latestVisit.node.right
			depth = latestVisit.depth // reset the depth counter as required for this level
		}

		if current != nil {
			depth++
		}
	}

	return visited
}

func extractNodesFromVisited(visited []*visitedNode) []*Node {
	nodes := make([]*Node, 0)
	for _, vnode := range visited {
		nodes = append(nodes, vnode.node)
	}

	return nodes
}
