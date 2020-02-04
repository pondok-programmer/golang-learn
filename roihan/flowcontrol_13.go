package roihan

import "fmt"

func flowcontrol13() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done!")
}
