package main

/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */

// @lc code=start
type node struct {
	isEnd    bool
	children map[byte]*node
}

// Trie type
type Trie struct {
	root *node
}

// Constructor Trie
/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		root: &node{
			children: map[byte]*node{},
		},
	}
}

// Insert func
/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	nd := t.root
	for i := 0; i < len(word); i++ {
		b := word[i]
		if _, ok := nd.children[b]; ok {
			nd = nd.children[b]
			continue
		} else {
			nd.children[b] = &node{
				children: map[byte]*node{},
			}
			nd = nd.children[b]
		}
	}
	nd.isEnd = true
}

// Search func
/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	nd := t.root
	for i := 0; i < len(word); i++ {
		b := word[i]
		if tmp, ok := nd.children[b]; ok {
			nd = tmp
		} else {
			return false
		}
	}
	return nd.isEnd
}

// StartsWith func
/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	nd := t.root
	for i := 0; i < len(prefix); i++ {
		b := prefix[i]
		if tmp, ok := nd.children[b]; ok {
			nd = tmp
		} else {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
