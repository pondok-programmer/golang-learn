package roihan

import "fmt"

var pos = []int{1, 2, 4, 8, 16, 32, 64, 128}

func moretype16() {
	for i, v := range pos {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
