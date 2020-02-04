package sidik

import "fmt"

func flow3()  {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	
	fmt.Println(sum)
}