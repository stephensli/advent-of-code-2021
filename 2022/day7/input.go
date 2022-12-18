package main

import (
	"strconv"
	"strings"
	"unicode"
)

// cd means change directory. This changes which directory is the current
// directory, but the specific result depends on the argument:

// - cd x moves in one level: it looks in the current directory for the directory
// named x and makes it the current directory.

// - cd .. moves out one level: it finds the directory that contains the current
// directory, then makes that directory the current directory.

// - cd / switches the current directory to the outermost directory, /.
func handleCD(directory string, root *Node, currentNode *Node) *Node {
	// If we are pointing to the current directly, no reason to handle a funky
	// iterative loop to the root. Just jump to the root.
	if directory == "/" {
		return root
	}

	// handle going up a single level
	if directory == ".." {
		return currentNode.parent
	}

	return currentNode.children[directory]
}

func handleFile(line string, _ *Node, currentNode *Node) {
	file := strings.Split(line, " ")
	size, _ := strconv.ParseInt(file[0], 10, 64)

	if _, ok := currentNode.children[file[1]]; !ok {
		fileNode := NewNode(currentNode, File, file[1], size)
		currentNode.children[file[1]] = fileNode
	}
}

// - dir xyz means that the current directory contains a directory named xyz.
func handleDir(directory string, _ *Node, currentNode *Node) {
	directoryNode := NewNode(currentNode, Directory, directory, 0)

	if _, ok := currentNode.children[directory]; !ok {
		currentNode.children[directory] = directoryNode
	}
}

func parseInput(lines []string) *Node {
	root := NewNode(nil, Directory, "/", 0)

	currentNode := root

	index := 0
	lines = lines[1:]

	for index != len(lines) {
		line := lines[index]

		if strings.HasPrefix(line, "$ cd") {
			currentNode = handleCD(strings.Split(line, "$ cd ")[1], root, currentNode)
			index += 1
			continue
		}

		if strings.HasPrefix(line, "$ ls") {
			index += 1
			continue
		}

		if strings.HasPrefix(line, "dir") {
			handleDir(strings.Split(line, " ")[1], root, currentNode)
			index += 1
			continue
		}

		if unicode.IsNumber(rune(line[0])) {
			handleFile(line, root, currentNode)
			index += 1
			continue
		}
	}

	// cache := map[string]int64{}
	// root.Print("", cache)

	return root
}
