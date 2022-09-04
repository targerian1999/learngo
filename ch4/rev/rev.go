package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	// fmt.Println(a)

	s := []int{0, 1, 2, 3, 4}
	reverse(s[:2])
	fmt.Printf("%d\n", s)
	reverse(s[2:])
	fmt.Printf("%d\n", s)
	reverse(s)
	fmt.Println(s)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
