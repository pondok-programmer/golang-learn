package sidik

import "fmt"

func type10() {
	s := []int{4,5,3,1,3,8,9}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[:1]
	fmt.Println(s)

}