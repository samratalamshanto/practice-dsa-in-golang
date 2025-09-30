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

func CreateSlice(len int, cap int) *[]int {
	var slice = make([]int, len, cap)
	return &slice
}
