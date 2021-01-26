package main

import "fmt"

func avg(arg ...float64) (avgAns float64) {
	var sum float64
	if len(arg) == 0 {
		return 0
	}

	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}
	avgAns = sum / float64(len(arg))
	return
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}

	// Print the output
	/*
		for i:=0; i<len(ar); i++ {
			var ar []int = []int{5,3,4,1,2,10,11,0}
			println(ar[i])
		}
	*/
}

// Write a function that takes a variable number of ints
// and print each integer on a separate line.
func varArgs(args ...int) {
	if len(args) == 0 {
		fmt.Println("No input")
	}
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}

func fibonacci(n int) []int {
	var fibb = make([]int, n)
	fibb[0] = 1
	fibb[1] = 1

	for i := 2; i < n; i++ {
		fibb[i] = fibb[i-1] + fibb[i-2]
	}
	return fibb
}

/*
Write a function that returns a function that performs a +2on integers.
Name the function plusTwo. You should then be able do the following:
p := plusTwo()
fmt.Printf("%v\n", p(2))
*/
func plusTwo() func(n int) int {
	//         --- return type -
	return func(n int) int { return n + 2 }
}

func plusX(x int) func(int) int {
	return func(n int) int { return n + x }
}

func maximum(n []int) (max int) {
	max = n[0]
	for i := 1; i < len(n); i++ {
		if n[i] > max {
			max = n[i]
		}
	}
	return
}

func Map(x func(int) int, ls []int) {
	for i := 0; i < len(ls); i++ {
		ls[i] = x(ls[i])
	}
}

// associated to the Map question
func doSomething(x int) int {
	return x * 2
}

func main() {
	StackMain()
}
