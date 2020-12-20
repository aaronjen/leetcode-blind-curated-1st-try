package main

/*
 * @lc app=leetcode id=211 lang=golang
 *
 * [211] Design Add and Search Words Data Structure
 */

// @lc code=start

type node struct {
	isEnd    bool
	children map[byte]*node
}

// WordDictionary type
type WordDictionary struct {
	root *node
}

// Constructor func
/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		root: &node{
			children: map[byte]*node{},
		},
	}
}

// AddWord func
func (dic *WordDictionary) AddWord(word string) {
	nd := dic.root
	for i := 0; i < len(word); i++ {
		b := word[i]
		if tmp, ok := nd.children[b]; ok {
			nd = tmp
		} else {
			tmp := &node{
				children: map[byte]*node{},
			}
			nd.children[b] = tmp
			nd = tmp
		}
	}
	nd.isEnd = true
}

// Search func
func (dic *WordDictionary) Search(word string) bool {
	return searhHelper(word, 0, dic.root)
}

func searhHelper(word string, ind int, nd *node) bool {
	if len(word) == ind {
		return nd.isEnd
	}

	b := word[ind]
	if b == '.' {
		for _, nextNd := range nd.children {
			if searhHelper(word, ind+1, nextNd) {
				return true
			}
		}
	} else {
		if nextNd, ok := nd.children[b]; ok {
			return searhHelper(word, ind+1, nextNd)
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end
