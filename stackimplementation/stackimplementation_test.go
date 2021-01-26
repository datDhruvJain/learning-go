package stackimplementation

import (
	"fmt"
	"testing"
)

func TestCreateStack(t *testing.T) {
	fmt.Println("\nCreating new Stack of size 10")
	s := new(Stack)
	s.New(10)
	fmt.Print("\nCreated new stack, len: ")
	fmt.Println(len(s.stk))
	fmt.Println()
}

func TestPush(t *testing.T) {
	s := new(Stack)
	s.New(5)
	s.PushStack(12)
	s.PushStack(3)
	s.PushStack(23)
	fmt.Printf("The stack after pushing elements is %v\n", s.stk)
}

// Demonstrating how Examples work in GOlang
func ExamplePushStack() {
	s := new(Stack)
	s.New(5)
	s.PushStack(12)
	s.PushStack(3)
	s.PushStack(23)
	fmt.Println(s.stk)
	// Output:
	// [23 3 12 0 0]
}
