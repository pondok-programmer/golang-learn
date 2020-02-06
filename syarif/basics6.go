package syarif

import "fmt"

func swap(a, b string) (string, string) {
	return a, b
}

func basics6() {
	a, b := "Hello", "World"
	fmt.Println(swap(a, b))
}
