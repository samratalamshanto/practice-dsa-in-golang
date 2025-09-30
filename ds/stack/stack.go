package main

import "fmt"

func main() {
	stack := []int{}
	stack = append(stack, 1)
	stack = append(stack, 2, 3, 4, 5) //add

	fmt.Println(stack)

	top := stack[len(stack)-1] //top
	fmt.Println(top)

	stack = stack[:len(stack)-1] //pop
	fmt.Println(stack)

}
