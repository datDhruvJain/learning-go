package main

import "fmt"

func main() {
	prob2()
}

// Create a loop with the for construct.
// Make it loop 10 times and print out the loop counter with the fmt package.
func prob1() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Println(i)
	}
}

// Rewrite the loop from 1 to use goto. The keyword for may not be used.
func prob2() {
	var i int = 0
loop:

	fmt.Println(i)
	i++
	if i < 10 {
		goto loop
	}
}
