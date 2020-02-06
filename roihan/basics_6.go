package roihan

import "fmt"

func swap(x, y string) (string, string)  {
	return x, y
}

func basics6() {
	a, b := swap("Hello", "World")
	fmt.Println(a,b)
}