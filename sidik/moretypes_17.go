package sidik

import "fmt"

func type17() {
	ppow := make([]int, 10)
	for i := range ppow {
		ppow[i] = 1 <<  uint(i) // 2**i

	}

	for _, value := range ppow {
		fmt.Printf("%d\n", value)
	}
}