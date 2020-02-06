package sidik

import "fmt"

func flow13()  {
	fmt.Println("counting!")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}