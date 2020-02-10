package roihan

import "fmt"

func moretype8() {
	names := [3]string{
		"Roihan",
		"Mishbahul",
		"Anam",
	}

	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
