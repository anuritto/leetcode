package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(initNode *Node) *Node {
	if initNode == nil {
		return nil
	}
	copies := make(map[*Node]*Node)
	return recursion(copies, initNode)
}

func recursion(copies map[*Node]*Node, node *Node) *Node {
	existed := copies[node]
	if existed != nil {
		return existed
	}

	copy := &Node{
		Val:       node.Val,
	}
	copies[node] = copy

	for i := 0; i < len(node.Neighbors); i++ {
		neighbourCopy := recursion(copies, node.Neighbors[i])
		copy.Neighbors = append(copy.Neighbors, neighbourCopy)
	}

	return copy
}

// https://leetcode.com/problems/clone-graph/