package main

import "fmt"

func main() {

	m := map[string]int{}
	m["abc"] = 1
	m["a"] = 2
	m["b"] = 3

	fmt.Println(m)

	delete(m, "a")
	fmt.Println(m)

	for key, val := range m {
		fmt.Printf("key=%v, value=%v \n", key, val)
	}

	val, ok := m["cyz"]
	fmt.Printf("value=%v, ok=%v \n", val, ok)

	val, ok = m["abc"]
	fmt.Printf("value=%v, ok=%v \n", val, ok)
}
