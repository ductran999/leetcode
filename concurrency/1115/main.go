package main

import (
	"fmt"
	"time"
)

type FooBar struct {
	n    int
	fooC chan struct{}
	barC chan struct{}
}

func NewFooBar(n int) *FooBar {
	fb := &FooBar{
		n:    n,
		fooC: make(chan struct{}, 1),
		barC: make(chan struct{}, 1),
	}
	fb.fooC <- struct{}{} // preload so Foo starts first
	return fb
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.fooC             // wait until allowed to print foo
		printFoo()            // print "foo"
		fb.barC <- struct{}{} // signal bar to run
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.barC             // wait until allowed to print bar
		printBar()            // print "bar"
		fb.fooC <- struct{}{} // signal foo to run
	}
}

func main() {
	f := NewFooBar(2)
	go f.Foo(func() {
		fmt.Printf("foo")
	})
	go f.Bar(func() {
		fmt.Print("bar")
	})
	time.Sleep(time.Second)
}
