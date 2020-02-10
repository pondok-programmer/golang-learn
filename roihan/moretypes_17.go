package roihan

import "fmt"

func moretype17() {
	por := make([]int, 10)
	for i := range por {
		por[i] = 1 << uint(i)
	}
	for _, value := range por {
		fmt.Printf("%d\n", value)
	}
}
