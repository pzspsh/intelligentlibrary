/*
@File   : prefixtree.go
@Author : pan
@Time   : 2023-05-25 09:41:32
*/
package main

func main() {

}

// 实现Trie(前缀树)
type Trie struct {
	next   [26]*Trie
	ending int
}

func Constructor() Trie {
	return Trie{
		next:   [26]*Trie{},
		ending: 0,
	}
}

func (t *Trie) Insert(word string) {
	temp := t
	for _, v := range word {
		value := v - 'a'
		if temp.next[value] == nil {
			temp.next[value] = &Trie{
				next:   [26]*Trie{},
				ending: 0,
			}
		}
		temp = temp.next[value]
	}
	temp.ending++
}

func (t *Trie) Search(word string) bool {
	temp := t
	for _, v := range word {
		value := v - 'a'
		if temp = temp.next[value]; temp == nil {
			return false
		}
	}
	if temp.ending > 0 {
		return true
	} else {
		return false
	}
}

func (t *Trie) StartsWith(prefix string) bool {
	temp := t
	for _, v := range prefix {
		value := v - 'a'
		if temp = temp.next[value]; temp == nil {
			return false
		}
	}
	return true
}
