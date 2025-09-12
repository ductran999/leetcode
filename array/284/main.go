package main

import "fmt"

type Iterator struct {
	arr []int
	idx int
}

func NewIterator(nums []int) *Iterator {
	return &Iterator{
		arr: nums,
		idx: -1,
	}
}

func (it *Iterator) hasNext() bool {
	return it.idx < len(it.arr)-1
}

func (it *Iterator) next() int {
	it.idx++
	return it.arr[it.idx]
}

type PeekingIterator struct {
	iter       *Iterator
	prevVal    int
	currVal    int
	stackEmpty bool // keep track last element is pop or not
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		iter:       iter,
		currVal:    iter.next(), // go to first element
		prevVal:    0,
		stackEmpty: false,
	}
}

func (pk *PeekingIterator) hasNext() bool {
	// the iterator is go to end but not pop the last element
	if !pk.iter.hasNext() && !pk.stackEmpty {
		return true
	}
	if !pk.iter.hasNext() && pk.stackEmpty {
		return false
	}
	return pk.iter.hasNext()
}

func (pk *PeekingIterator) next() int {
	pk.prevVal = pk.currVal
	pk.currVal = pk.iter.next()
	if pk.currVal < 0 {
		// the pointer is point out of list
		pk.stackEmpty = true
	}
	return pk.prevVal
}

func (pk *PeekingIterator) peek() int {
	return pk.currVal
}

func main() {
	nums := []int{1, 2, 3}
	itr := NewIterator(nums)

	pk := Constructor(itr)
	fmt.Println(pk.next())
	fmt.Println(pk.hasNext())
	fmt.Println(pk.peek())
}
