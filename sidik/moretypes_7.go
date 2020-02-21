package sidik

import "fmt"

func type7() {
	prime := [6]int{12,3,42,4,5,3}

	var s []int = prime[1:4]
	fmt.Println(s)
}