package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)
	s = make([]int, 10, 20)        // array interno do slice definido para 20
	fmt.Println(s, len(s), cap(s)) // cap() capacidade do array interno do slice
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}
