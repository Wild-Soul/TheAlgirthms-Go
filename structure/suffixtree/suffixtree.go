// Package suffixtree provides an implementation of Suffix Tree data structures in golang.
// Suffix Trees can be considered as a compressed version of Trie, where each node doesn't just contains one character but can contain a substring.
// For this implementation we are going to use Ukkonen's algorithm.
//
// Wikipedia: https://en.wikipedia.org/wiki/Suffix_tree
// Cp-algorithms: https://cp-algorithms.com/string/suffix-tree-ukkonen.html
package suffixtree

import (
	"fmt"
)

type Node struct {
	children   map[byte]*Node
	suffixLink *Node
	start, end *int
	id         int
}

type SuffixTree struct {
	root      *Node
	text      string
	nodeCount int
}

func NewNode(start, end *int) *Node {
	return &Node{
		children:   make(map[byte]*Node),
		suffixLink: nil,
		start:      start,
		end:        end,
		id:         -1,
	}
}

func NewSuffixTree(text string) *SuffixTree {
	root := NewNode(nil, nil)
	return &SuffixTree{
		root:      root,
		text:      text + "$", // Add a terminator character
		nodeCount: 0,
	}
}

func (st *SuffixTree) Build() {
	n := len(st.text)
	rootEnd := -1
	st.root.end = &rootEnd
	activeNode := st.root
	activeEdge := -1
	activeLength := 0
	remainingSuffixCount := 0

	for i := 0; i < n; i++ {
		remainingSuffixCount++
		lastNewNode := (*Node)(nil)

		for remainingSuffixCount > 0 {
			if activeLength == 0 {
				activeEdge = i
			}

			if _, exists := activeNode.children[st.text[activeEdge]]; !exists {
				st.nodeCount++
				end := i
				activeNode.children[st.text[activeEdge]] = NewNode(&i, &end)

				if lastNewNode != nil {
					lastNewNode.suffixLink = activeNode
					lastNewNode = nil
				}
			} else {
				next := activeNode.children[st.text[activeEdge]]
				if st.walkDown(&activeNode, &i, &activeEdge, &activeLength) {
					continue
				}

				if st.text[*next.start+activeLength] == st.text[i] {
					if lastNewNode != nil && activeNode != st.root {
						lastNewNode.suffixLink = activeNode
						lastNewNode = nil
					}
					activeLength++
					break
				}

				splitEnd := *next.start + activeLength - 1
				split := NewNode(next.start, &splitEnd)
				st.nodeCount++
				activeNode.children[st.text[activeEdge]] = split

				end := i
				split.children[st.text[i]] = NewNode(&i, &end)
				st.nodeCount++
				next.start = new(int)
				*next.start = *split.end + 1
				split.children[st.text[*next.start]] = next

				if lastNewNode != nil {
					lastNewNode.suffixLink = split
				}
				lastNewNode = split
			}

			remainingSuffixCount--
			if activeNode == st.root && activeLength > 0 {
				activeLength--
				activeEdge = i - remainingSuffixCount + 1
			} else if activeNode != st.root {
				activeNode = activeNode.suffixLink
			}
		}
	}
}

func (st *SuffixTree) walkDown(activeNode **Node, i, activeEdge, activeLength *int) bool {
	node := (*activeNode).children[st.text[*activeEdge]]
	length := *node.end - *node.start + 1
	if *activeLength < length {
		return false
	}
	*activeEdge += length
	*activeLength -= length
	*activeNode = node
	return true
}

func (st *SuffixTree) DFS(node *Node, depth int) {
	if node == nil {
		return
	}

	if node != st.root {
		fmt.Printf("Edge: %s, Node ID: %d\n", st.text[*node.start:*node.end+1], node.id)
	}

	for _, child := range node.children {
		st.DFS(child, depth+1)
	}
}

func (st *SuffixTree) Contains(pattern string) bool {
	node := st.root
	i := 0
	for i < len(pattern) {
		if child, exists := node.children[pattern[i]]; exists {
			edge := st.text[*child.start : *child.end+1]
			if len(pattern[i:]) >= len(edge) && pattern[i:i+len(edge)] == edge {
				i += len(edge)
				node = child
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
