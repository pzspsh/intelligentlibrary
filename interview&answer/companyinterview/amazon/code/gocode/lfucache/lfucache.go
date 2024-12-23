/*
@File   : lfucache.go
@Author : pan
@Time   : 2023-05-22 09:52:29
*/
package main

import "container/list"

func main() {

}

// LFU缓存
type Node struct {
	key   int
	value int
	count int
}

type LFUCache struct {
	cap     int
	minFreq int
	kv      map[int]*list.Element
	fk      map[int]*list.List
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:     capacity,
		minFreq: 0,
		kv:      make(map[int]*list.Element),
		fk:      make(map[int]*list.List),
	}
}

func (l *LFUCache) Get(key int) int {
	node, ok := l.kv[key]
	if !ok {
		return -1
	}
	return l.increaseFreq(node)
}

func (l *LFUCache) Put(key int, value int) {
	data, ok := l.kv[key]
	if ok {
		node := data.Value.(*Node)
		node.value = value
		l.increaseFreq(data)
		return
	}
	if l.cap == len(l.kv) {
		cur, ok := l.fk[l.minFreq]
		if !ok {
			return
		}
		deleteKey := cur.Front()
		cur.Remove(deleteKey)
		delete(l.kv, deleteKey.Value.(*Node).key)
	}
	temp := &Node{
		key:   key,
		value: value,
		count: 1,
	}
	if _, ok := l.fk[1]; !ok {
		l.fk[1] = list.New()
	}
	res := l.fk[1].PushBack(temp)
	l.kv[key] = res
	l.minFreq = 1
}

func (l *LFUCache) increaseFreq(data *list.Element) int {
	node := data.Value.(*Node)
	cur, ok := l.fk[node.count]
	if !ok {
		return -1
	}
	cur.Remove(data)
	if cur.Len() == 0 && l.minFreq == node.count {
		l.minFreq++
	}
	node.count++
	if l.fk[node.count] == nil {
		l.fk[node.count] = list.New()
	}
	res := l.fk[node.count].PushBack(node)
	l.kv[node.key] = res
	return node.value
}
