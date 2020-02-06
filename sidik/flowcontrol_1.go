package sidik

import "fmt"

func flow1()  {
	sum := 1

	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}