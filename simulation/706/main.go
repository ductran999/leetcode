package main

import "fmt"

type Node struct {
	key, value int
	next       *Node
}

type MyHashMap struct {
	buckets []*Node
	size    int
}

func Constructor() MyHashMap {
	return MyHashMap{
		buckets: make([]*Node, 1000), // pick a base size (prime number is good too)
		size:    1000,
	}
}

func (hm *MyHashMap) hash(key int) int {
	return key % hm.size
}

func (hm *MyHashMap) Put(key int, value int) {
	idx := hm.hash(key)
	head := hm.buckets[idx]

	// Search for existing key
	for node := head; node != nil; node = node.next {
		if node.key == key {
			node.value = value
			return
		}
	}

	// Insert new node at head
	newNode := &Node{key: key, value: value, next: head}
	hm.buckets[idx] = newNode
}

func (hm *MyHashMap) Get(key int) int {
	idx := hm.hash(key)
	for node := hm.buckets[idx]; node != nil; node = node.next {
		if node.key == key {
			return node.value
		}
	}
	return -1
}

func (hm *MyHashMap) Remove(key int) {
	idx := hm.hash(key)
	head := hm.buckets[idx]
	var prev *Node

	for node := head; node != nil; node = node.next {
		if node.key == key {
			if prev == nil {
				hm.buckets[idx] = node.next
			} else {
				prev.next = node.next
			}
			return
		}
		prev = node
	}
}

func main() {
	m := Constructor()
	m.Put(2, 4)
	fmt.Println(m.Get(2))
}
