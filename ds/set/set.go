package main

import "fmt"

type Set map[string]struct{}

func (s Set) Append(val string) {
	//an empty struct that occupies 0 bytes of memory
	s[val] = struct{}{} //struct{} for struct type and {} --> for empty obj
}

func (s Set) Delete(key string) {
	delete(s, key)
}

func (s Set) Contains(key string) bool {
	_, ok := s[key]
	return ok
}

func (s Set) Len() int {
	return len(s)
}

func (s Set) Union(other Set) Set {
	result := make(Set)
	for k := range s {
		result.Append(k)
	}

	for k := range other {
		result.Append(k)
	}

	return result
}

func (s Set) Intersection(other Set) Set {
	result := make(Set)
	for k := range s {
		if other.Contains(k) {
			result.Append(k)
		}
	}

	return result
}

func (s Set) Difference(other Set) Set {
	result := make(Set)
	for k := range s {
		if !other.Contains(k) {
			result.Append(k)
		}
	}

	return result
}

func main() {
	s := make(Set)
	s.Append("10")
	s.Append("20")
	s.Append("30")

	fmt.Println(s)

	fmt.Println(s.Contains("10"))
	fmt.Println(s.Contains("30"))

	s.Delete("30")
	fmt.Println(s.Contains("30"))
	fmt.Println(s)

	other := make(Set)
	other.Append("a")
	other.Append("b")
	other.Append("10")
	fmt.Println("Other Set:")
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
