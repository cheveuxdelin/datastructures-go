package trie

import "github.com/cheveuxdelin/datastructures-go/stack"

type node struct {
	value    rune
	isWord   bool
	children map[rune]*node
}

type Trie struct {
	root node
}

type stackStruct struct {
	node *node
	word string
}

func (t *Trie) Push(s string) {
	if t.root.children == nil {
		t.root.children = make(map[rune]*node)
	}
	var current *node = &t.root
	for _, c := range s {
		// check if node exists
		if _, ok := current.children[c]; !ok {
			current.children[c] = &node{value: c, isWord: false, children: make(map[rune]*node)}
		}
		current = current.children[c]
	}
	current.isWord = true
}

func (t *Trie) IsWord(s string) bool {
	if t.root.children == nil {
		t.root.children = make(map[rune]*node)
	}
	var current *node = &t.root
	for _, c := range s {
		// check if node exists
		if _, ok := current.children[c]; !ok {
			return false
		}
		current = current.children[c]
	}
	return current.isWord
}

func (t *Trie) GetWords(s string) []string {
	var rtn []string
	if t.root.children == nil {
		return rtn
	}
	var current *node = &t.root
	for _, c := range s {
		if _, ok := current.children[c]; !ok {
			return rtn
		} else {
			current = current.children[c]
		}
	}

	q := stack.Stack[stackStruct]{}
	q.Push(stackStruct{current, s})

	for q.Size() != 0 {
		current, _ := q.Pop()
		if current.node.isWord {
			rtn = append(rtn, current.word)
		}
		for c, node := range current.node.children {
			q.Push(stackStruct{node: node, word: current.word + string(c)})
		}
	}

	return rtn
}
