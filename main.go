package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	var name string
	fmt.Scanln(&name)
	fmt.Println(name)
}
