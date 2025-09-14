package main

import "fmt"

type MyHashSet struct {
	List [1000001]bool
}

func Constructor() MyHashSet {
	return MyHashSet{}
}

func (hs *MyHashSet) Add(key int) {
	hs.List[key] = true
}

func (hs *MyHashSet) Remove(key int) {
	hs.List[key] = false
}

func (hs *MyHashSet) Contains(key int) bool {
	return hs.List[key]
}

func main() {
	m := Constructor()
	m.Add(2)
	fmt.Println(m.Contains(2))
}
