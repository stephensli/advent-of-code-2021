package main

import (
	"encoding/base64"
	"fmt"
	"sort"

	"github.com/stephensli/aoc/helpers/queue"
)

type Kind int

const (
	Directory Kind = iota
	File      Kind = iota
)

type Node struct {
	kind     Kind
	parent   *Node
	children map[string]*Node
	name     string
	size     int64
	key      string
}

func (n *Node) Size(cache map[string]int64) int64 {
	// cache has to be based on all the parent names, otherwise you will end up
	// using an existing cache for a duplicated file name.
	if val, ok := cache[n.key]; ok {
		return val
	}

	if n.kind == File {
		return n.size
	}

	var total int64
	for _, node := range n.children {
		total += node.Size(cache)
	}

	cache[n.key] = total
	return total
}

func (n *Node) Print(prefix string, cache map[string]int64) {
	switch n.kind {
	case Directory:
		fmt.Printf("%s - %s (dir, size=%d)\n", prefix, n.name, n.Size(cache))
	case File:
		fmt.Printf("%s - %s (file, size=%d)\n", prefix, n.name, n.size)
	}

	for _, node := range n.children {
		node.Print(prefix+" ", cache)
	}
}

func NewNode(parent *Node, kind Kind, name string, size int64) *Node {
	var key string

	if parent != nil {
		key = base64.StdEncoding.EncodeToString([]byte(parent.key + name))
	} else {
		key = base64.StdEncoding.EncodeToString([]byte(name))
	}

	return &Node{
		children: map[string]*Node{},
		kind:     kind,
		name:     name,
		parent:   parent,
		size:     size,
		key:      key,
	}
}

type NodeSize struct {
	name string
	kind Kind
	size int64
}

func directorySizes(root *Node) []NodeSize {
	var sizes []NodeSize
	cache := map[string]int64{} // will contain more than just directories

	if root.kind == Directory {
		sizes = append(sizes, NodeSize{
			name: root.name,
			kind: root.kind,
			size: root.Size(cache),
		})
	}

	todo := queue.Stack[*Node]{}
	for _, node := range root.children {
		todo.Push(node)
	}

	for todo.Len() != 0 {
		node := todo.Pop()

		if node.kind == Directory {
			sizes = append(sizes, NodeSize{
				name: node.name,
				kind: node.kind,
				size: node.Size(cache),
			})
		}

		for _, n := range node.children {
			todo.Push(n)
		}
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i].size > sizes[j].size
	})

	return sizes
}
