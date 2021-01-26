package stackimplementation

// Stack : declares a stack
type Stack struct {
	stk []int
}

// New : Create new stack
func (s *Stack) New(len int) {
	var st = make([]int, len)
	s.stk = st
}

// PushStack : Push an element onto the stack
func (s *Stack) PushStack(e int) {
	for i := len(s.stk) - 1; i > 0; i-- {
		s.stk[i] = s.stk[i-1]
	}
	s.stk[0] = e
}

// PopStack : Pop the first element from the stack
func PopStack(stack []int) (e int) {
	e = stack[0]
	for i := 0; i < len(stack)-1; i++ {
		stack[i] = stack[i+1]
	}
	stack[len(stack)-1] = 0
	return
}
