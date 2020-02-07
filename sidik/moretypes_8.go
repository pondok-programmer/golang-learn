package sidik

import "fmt"

func type8() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	b := names[1:3]
	a := names[0:2]

	fmt.Println(b, a)

	b[0] = "XXXx"
	fmt.Println(b, a)
	fmt.Println(names)

}