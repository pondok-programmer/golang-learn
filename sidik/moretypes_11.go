package sidik

import "fmt"

func type11() {
	s := []int{5, 2, 6, 7, 8, 9, 10}
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)
}

func printSlice(s []int)  {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}