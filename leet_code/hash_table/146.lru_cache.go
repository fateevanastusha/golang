package main

import (
	"fmt"
)

/*
проблемы, с которыми мы сталкиваемся при формировании кеша
1) invalidation (когда в бд лежат обновленные данные, а в кеше устаревшие. забыли обновить данные в кеше тоже)
2) прогрев (процесс предварительной загрузки часто используемых данных при инициализации)
3) политика удаления (когда память заканчивается нужно кого-то удалить. например, lru (last recently used))

LRU - всегда linked list и hash


в начале самые свежие
в конце самые старые.
удаляем из конца, добавляем в начало.
*/

type Node struct {
	Prev, Next *Node
	Key, Value int
}

func newNode(key, val int) *Node {
	return &Node{
		Key: key, Value: val,
	}
}

type LRUCache struct {
	//вместимость
	capacity int
	//значения со ссылкой на ноду, чтобы можно было из середины ее переместить куда-то
	values map[int]*Node
	//первая и последняя
	head, tail *Node
}

func newLRUCache(head, tail *Node, capacity int) LRUCache {
	return LRUCache{
		head:     head,
		tail:     tail,
		capacity: capacity,
		values:   make(map[int]*Node),
	}
}

func Constructor(capacity int) LRUCache {
	h, t := newNode(0, 0), newNode(0, 0)
	h.Next = t
	t.Prev = h
	return newLRUCache(h, t, capacity)
}

// time - O(1)
func (this *LRUCache) Get(key int) int {
	if node, ok := this.values[key]; ok {
		//удаляем из середины/конца
		this.remove(node)
		//добавляем в начало
		this.insert(node)
		return node.Value
	}

	return -1

}

// time - O(1)
func (this *LRUCache) Put(key int, value int) {

	if v, ok := this.values[key]; ok {
		//delete if already have
		this.remove(v)
	}
	if len(this.values) == this.capacity {
		//удаляем самую старую
		this.remove(this.tail.Prev)
	}

	//добавляем в начало
	this.insert(newNode(key, value))

}

func (this *LRUCache) insert(node *Node) {
	this.values[node.Key] = node
	next := this.head.Next
	this.head.Next = node
	node.Prev = this.head
	next.Prev = node
	node.Next = next
}
func (this *LRUCache) remove(node *Node) {
	delete(this.values, node.Key)
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}
func main() {

	//1 3
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))

}
