package sidik

import "fmt"

func flow12()  {
	defer fmt.Println("World")
	defer fmt.Println("Hello")

	fmt.Println("Begin")
}