package main

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return true
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	return 1
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

func (this *PeekingIterator) hasNext() bool {
	// the iterator is go to end but not pop the last element
	if !this.iter.hasNext() && !this.stackEmpty {
		return true
	}
	if !this.iter.hasNext() && this.stackEmpty {
		return false
	}
	return this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	this.prevVal = this.currVal
	this.currVal = this.iter.next()
	if this.currVal < 0 {
		// the pointer is point out of list
		this.stackEmpty = true
	}
	return this.prevVal
}

func (this *PeekingIterator) peek() int {
	return this.currVal
}
