package main

import "fmt"

func Stack(len int) []int {
	var st = make([]int, len)
	return st
}

func PushStack(e int, stack []int) {
	for i := len(stack) - 1; i > 0; i-- {
		stack[i] = stack[i-1]
	}
	stack[0] = e
}

func PopStack(stack []int) (e int) {
	e = stack[0]
	for i := 0; i < len(stack)-1; i++ {
		stack[i] = stack[i+1]
	}
	stack[len(stack)-1] = 0
	return
}

func StackMain() {
	ls := Stack(5)
	PushStack(9, ls)
	PushStack(1, ls)
	PushStack(2, ls)
	PushStack(3, ls)
	PushStack(4, ls)
	PushStack(5, ls)
	fmt.Println(ls)
	fmt.Println(PopStack(ls))
	fmt.Println(ls)
	fmt.Println(PopStack(ls))
	fmt.Println(PopStack(ls))
	fmt.Println(PopStack(ls))
	fmt.Println(ls)
	fmt.Println(PopStack(ls))
}
