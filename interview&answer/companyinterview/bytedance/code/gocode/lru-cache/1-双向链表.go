package main

import "fmt"

func main() {
	node := Constructor(3)
	node.Put(2, 1)
	fmt.Println(node)
}

// LRU缓存机制
type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	cap    int
	header *Node
	tail   *Node
	m      map[int]*Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		cap:    capacity,
		header: &Node{},
		tail:   &Node{},
		m:      make(map[int]*Node, capacity),
	}
	cache.header.next = cache.tail
	cache.tail.prev = cache.header
	return cache
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.m[key]; ok {
		l.remove(node)
		l.putHead(node)
		return node.value
	}
	return -1
}

func (l *LRUCache) Put(key int, value int) {
	if node, ok := l.m[key]; ok {
		node.value = value
		l.remove(node)
		l.putHead(node)
		return
	}
	if l.cap <= len(l.m) {
		// 删除尾部
		deleteKey := l.tail.prev.key
		l.remove(l.tail.prev)
		delete(l.m, deleteKey)
	}
	// 插入到头部
	newNode := &Node{key: key, value: value}
	l.putHead(newNode)
	l.m[key] = newNode
}

// 删除尾部节点
func (l *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 插入头部
func (l *LRUCache) putHead(node *Node) {
	next := l.header.next
	l.header.next = node
	node.next = next
	next.prev = node
	node.prev = l.header
}
