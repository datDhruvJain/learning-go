package stackimplementation

import (
	"fmt"
	"testing"
)

/*func TestStack(t *testing.T) {

	ls := Stack(5)
	if len(ls) != 5 {
		t.Log("The test array created is not correct")
	}

	PushStack(9, ls)
	PushStack(1, ls)
	PushStack(2, ls)
	PushStack(3, ls)
	PushStack(4, ls)
	PushStack(5, ls)

	if PopStack(ls) != 5 {
		t.Log("Pop() return value wrong (1st location)")
	}
	fmt.Println(ls)
	fmt.Println(PopStack(ls))
	fmt.Println(PopStack(ls))
	fmt.Println(PopStack(ls))
	fmt.Println(ls)
	fmt.Println(PopStack(ls))
}
*/

func TestCreateStack(t *testing.T) {
	s := new(Stack)
	s.New(10)
	fmt.Println(len(s.stk))
}

func TestPush(t *testing.T) {
	s := new(Stack)
	s.New(5)
	s.PushStack(12)
	fmt.Println(s.stk)
}
