package ds

import "fmt"

func FixedArrSize5() *[5]int {
	var arr = [...]int{1, 2, 3, 4, 5}

	fmt.Println(arr)
	return &arr
}

func SliceArr() *[]int {
	var arr []int = []int{1, 2, 3, 4, 5}

	fmt.Println(arr)
	return &arr
}
