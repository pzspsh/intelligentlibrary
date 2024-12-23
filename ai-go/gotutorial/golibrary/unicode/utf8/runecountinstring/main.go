/*
@File   : main.go
@Author : pan
@Time   : 2023-12-04 10:13:38
*/
package main

import (
	"fmt"
	"unicode/utf8"
)

// 敏感词过滤
type Trie struct {
	child map[rune]*Trie
	word  string
}

// 插入
func (trie *Trie) insert(word string) *Trie {
	cur := trie
	for _, v := range []rune(word) {
		if _, ok := cur.child[v]; !ok {
			newTrie := NewTrie()
			cur.child[v] = newTrie
		}
		cur = cur.child[v]
	}
	cur.word = word
	return trie
}

// 过滤
func (trie *Trie) filerKeyWords(word string) string {
	cur := trie
	for i, v := range []rune(word) {
		if _, ok := cur.child[v]; ok {
			cur = cur.child[v]
			if cur.word != "" {
				word = replaceStr(word, "*", i-utf8.RuneCountInString(cur.word)+1, i)
				cur = trie
			}
		} else {
			cur = trie
		}
	}
	return word
}
func replaceStr(word string, replace string, left, right int) string {
	str := ""
	for i, v := range []rune(word) {
		if i >= left && i <= right {
			str += replace
		} else {
			str += string(v)
		}
	}
	return str
}
func NewTrie() *Trie {
	return &Trie{
		word:  "",
		child: make(map[rune]*Trie, 0),
	}
}
func main() {
	trie := NewTrie()
	trie.insert("sb").insert("狗日").insert("cnm").insert("狗日的")
	fmt.Println(trie.filerKeyWords("狗日，你就是个狗日的，我要cnm，你个sb，嘿嘿"))
}
