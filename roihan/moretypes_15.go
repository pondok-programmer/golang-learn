package roihan

import "fmt"

func moretype15() {
	var s []int
	Printslice(s)

	s = append(s, 0)
	Printslice(s)

	s = append(s, 1)
	Printslice(s)

	s = append(s, 2, 3, 4)
	Printslice(s)
}

func Printslice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
