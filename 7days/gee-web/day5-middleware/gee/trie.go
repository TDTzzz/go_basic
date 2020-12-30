package gee

import (
	"fmt"
	"log"
	"strings"
)

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.isWild)
}

func (n *node) insert(pattern string, parts []string, height int) {
	log.Println("node:", n)

	if len(parts) == height {
		n.pattern = pattern
		log.Println("node22:", n)
		return
	}
	//当前part
	part := parts[height]
	child := n.matchChild(part)

	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == '*' || part[0] == ':',
		}
		n.children = append(n.children, child)
	}
	log.Println("pattern:", pattern)
	log.Println("parts:", parts)
	log.Println("height:", height)
	log.Println("child:", child)
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			//todo 需要理解
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	log.Println("children:", children, "part:", part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}

func (n *node) matchChild(part string) *node {

	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
