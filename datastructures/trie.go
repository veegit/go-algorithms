package datastructures

import (
	"fmt"
)

//TrieNode node
type TrieNode struct {
	children     []*TrieNode
	isWordEnding bool
	wordTillNow  string
	value        byte
}

func (t *TrieNode) findChild(b byte) *TrieNode {
	for _, tn := range t.children {
		if tn.value == b {
			return tn
		}
	}
	return nil
}

//Trie trie
type Trie struct {
	root *TrieNode
}

//NewTrie new
func NewTrie() *Trie {
	trie := new(Trie)
	trieNode := TrieNode{value: "_"[0], wordTillNow: ""}
	trie.root = &trieNode
	return trie
}

//Insert insert
func (t *Trie) Insert(s string) {
	var current = t.root
	var node *TrieNode
	for i := range s {
		node = current.findChild(s[i])
		if node == nil {
			node = &TrieNode{value: s[i], wordTillNow: current.wordTillNow + string(s[i])}
			current.children = append(current.children, node)
		}
		current = node
		current.isWordEnding = i == len(s)-1
	}
}

//Search search
func (t *Trie) Search(s string) bool {
	var current = t.root
	for i := range s {
		var node = current.findChild(s[i])
		if node == nil {
			return false
		}
		current = node
	}
	return current != nil && current.isWordEnding
}

//Display di
func (t *Trie) Display() {
	current := t.root
	q := []*TrieNode{}
	q = append(q, current)
	for len(q) != 0 {
		current, q = q[0], q[1:]
		fmt.Printf("word: %s ", current.wordTillNow)
		q = append(q, current.children...)
	}
}

//Autocomplete comp
func (t *Trie) Autocomplete(s string) []string {
	var current = t.root
	for i := range s {
		var node = current.findChild(s[i])
		if node == nil {
			break
		}
		current = node
	}
	list := []string{}
	q := []*TrieNode{}
	q = append(q, current)
	for len(q) != 0 {
		current, q = q[0], q[1:]
		if current.isWordEnding {
			list = append(list, current.wordTillNow)
		}
		q = append(q, current.children...)
	}
	return list
}
