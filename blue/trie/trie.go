package trie

import (
	"strings"
)

type Node struct {
	Parttern string
	Parts    string
	Children []*Node
	IsWild   bool
}

// for create route, return the match child Node
func (n *Node) MatchChild(part string) *Node {
	for _, item := range n.Children {
		if item.Parts == part || n.IsWild {
			return item
		}
	}
	return nil
}

// for get route, return the list of children Node
func (n *Node) MatchChildren(parts string) []*Node {
	children := []*Node{}
	for _, item := range n.Children {
		if item.Parts == parts || item.IsWild {
			children = append(children, item)
		}
	}
	return children
}

// for create route
func (n *Node) Insert(pattern string, parts []string, height int) {
	// the last part
	if height == len(parts) {
		n.Parttern = pattern
		return
	}

	// match current part
	part := parts[height]
	child := n.MatchChild(part)

	// not existed current part, create a new one
	if child == nil {
		child = &Node{Parts: part, IsWild: part[0] == ':' || part[0] == '*'}
		n.Children = append(n.Children, child)
	}

	// continue enter the next part
	child.Insert(pattern, parts, height+1)
}

// for get route
func (n *Node) Search(parts []string, height int) *Node {
	if len(parts) == height || strings.HasPrefix(n.Parts, "*") {
		if n.Parttern == "" {
			return nil
		}
		return n
	}

	part := parts[height]

	// 获取当前节点的子节点
	children := n.MatchChildren(part)
	for _, child := range children {
		result := child.Search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
