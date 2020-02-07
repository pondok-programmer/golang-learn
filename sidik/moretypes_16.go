package sidik

import "fmt"

var poq = []int{1, 2, 4, 8, 16, 32, 64, 128}

func type16() {
	for i, v := range poq {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}