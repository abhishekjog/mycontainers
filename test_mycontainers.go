package main

import (
	"fmt"
	"mycontainers/mylist"
	"mycontainers/myoperations"
)

func main() {
	var ops myoperations.Operations
	ops = &mylist.SingleLinkedList{Head: nil}
	_ = ops.Push(10)
	_ = ops.Push(20)
	_ = ops.Push(30)

	r, t := ops.Pop()
	fmt.Println(r, t)
	r, t = ops.Pop()
	fmt.Println(r, t)
	r, t = ops.Pop()
	fmt.Println(r, t)
	r, t = ops.Pop()
	fmt.Println(r, t)
	r, t = ops.Pop()
	fmt.Println(r, t)
}
