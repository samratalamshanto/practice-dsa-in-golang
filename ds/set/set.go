package main

import "fmt"

type Set map[string]struct{}

func (s Set) Append(val string) {
	s[val] = struct{}{}
}

func (s Set) Delete(key string) {
	delete(s, key)
}

func (s Set) Contains(key string) bool {
	_, ok := s[key]
	return ok
}

func main() {
	s := make(Set)
	s.Append("10")
	s.Append("20")

	fmt.Println(s)

	fmt.Println(s.Contains("10"))
	fmt.Println(s.Contains("30"))

	s.Delete("10")
	fmt.Println(s.Contains("10"))
	fmt.Println(s)
}
