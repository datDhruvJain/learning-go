# learning-go

And logging my journey to this awesome language

Using this source to learn stuff quickly: https://miek.nl/go

By the way GO is a pass by value language. Arrays are also Passed by value, only the slices are pass by reference (till now)

for go getting with a private repo:
```shell
env GIT_TERMINAL_PROMPT=1 go get github.com/examplesite/myprivaterepo
```

<details markdown = "1"> 
<summary>Anonymous functions in go aka fuction as values</summary>

```go
import "fmt"

func main() {
	a := func() { 1
		fmt.Println("Hello")
	} 2
	a() 3
}
```
</details>
<details markdown = "1">
<summary> Functions and Callbacks</summary>

Because functions are values they are easy to pass to functions, from where they can be used as callbacks. First define a function that does “something” with an integer value:

```go
func printit(x int) {
    fmt.Printf("%v\n", x)
}
```
This function does not return a value and just prints its argument. The signature of this function is: func printit(int), or without the function name: func(int). To create a new function that uses this one as a callback we need to use this signature:
```go
func callback(y int, f func(int)) {
    f(y)
}
```
Here we create a new function that takes two parameters: y int, i.e. just an int and f func(int), i.e. a function that takes an int and returns nothing. The parameter f is the variable holding that function. It can be used as any other function, and we execute the function on line 2 with the parameter y: f(y)
</details>