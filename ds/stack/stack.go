package main

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) Append(val int) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() {
	if len(s.data) == 0 {
		return
	}
	s.data = s.data[:len(s.data)-1]
}

func (s *Stack) Top() int {
	if len(s.data) == 0 {
		return -1
	}
	return s.data[len(s.data)-1]
}

func main() {
	stack := &Stack{}
	fmt.Println(stack)

	stack.Append(2)
	stack.Append(3)
	stack.Append(4)
	stack.Append(5)
	fmt.Println(stack)

	top := stack.Top()
	fmt.Println(top)

	stack.Pop()
	fmt.Println(stack.data)

}
