package main

import "fmt"

func main() {
	fizzBuzz()
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

// Rewrite the loop again so that it fills an array
// and then prints that array to the screen.

func prob3() {
	i := 0
	var arr [10]int
	// Arrays are decalred as var arrayName [size]type

	for i = 0; i < len(arr); i++ {
		arr[i] = i
	}

	fmt.Println(arr)
}

// Write code to calculate the average of a float64 slice.
func prob4() {
	var avg float64 = 0
	var arr [10]float64 = [...]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s1 := arr[3:8]
	for i := 0; i < len(s1); {
		avg += s1[i]
		i++
	}
	avg = avg / float64(len(s1))
	fmt.Println(avg)
}

// Write a program that prints the numbers from 1 to 100.
// But for multiples of three print, “Fizz” instead of the number,
//  and for multiples of five, print “Buzz”.
// For numbers which are multiples of both three and five, print “FizzBuzz”.

func fizzBuzz() {
	for i := 0; i < 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")

		case i%3 == 0:
			fmt.Println("Fizz")

		case i%5 == 0:
			fmt.Println("Buzz")

		default:
			fmt.Println(i)

		}
	}
}
