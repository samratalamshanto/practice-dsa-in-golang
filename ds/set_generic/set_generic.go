package main

import "fmt"

type Set[T comparable] map[T]struct{}

func (s Set[T]) Append(val T) {
	//an empty struct that occupies 0 bytes of memory
	//struct{} for struct type(Empty) and {} --> initialize empty obj
	//type EmptyStuct struct{} and s[val]= EmptyStruct{}
	s[val] = struct{}{}
}

func (s Set[T]) Delete(key T) {
	delete(s, key)
}

func (s Set[T]) Contains(key T) bool {
	_, ok := s[key]
	return ok
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	result := make(Set[T])
	for k := range s {
		result.Append(k)
	}

	for k := range other {
		result.Append(k)
	}

	return result
}

func (s Set[T]) Intersection(other Set[T]) Set[T] {
	result := make(Set[T])
	for k := range s {
		if other.Contains(k) {
			result.Append(k)
		}
	}

	return result
}

func (s Set[T]) Difference(other Set[T]) Set[T] {
	result := make(Set[T])
	for k := range s {
		if !other.Contains(k) {
			result.Append(k)
		}
	}

	return result
}

func main() {
	s := make(Set[string]) //string type set create
	s.Append("10")
	s.Append("20")
	s.Append("30")

	fmt.Println(s)

	fmt.Println(s.Contains("10"))
	fmt.Println(s.Contains("30"))

	s.Delete("30")
	fmt.Println(s.Contains("30"))
	fmt.Println(s)

	other := make(Set[string]) //string type set create
	other.Append("a")
	other.Append("b")
	other.Append("10")
	fmt.Println("Other Set[T]:")
	fmt.Println(other)
	fmt.Println(other.Len())

	resultUnion := s.Union(other)
	fmt.Println("Union:")
	fmt.Println(resultUnion)

	resultIntersection := s.Intersection(other)
	fmt.Println("Intersection:")
	fmt.Println(resultIntersection)

	resultDifference := s.Difference(other)
	fmt.Println("Difference:")
	fmt.Println(resultDifference)
}
