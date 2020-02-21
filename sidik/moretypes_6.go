package sidik

import "fmt"

func type6() {
	var s [3]string
	s[0] = "sidik"
	s[1] = "sudah"
	s[2] = "go"

	fmt.Println(s[0], s[2])
	fmt.Println(s)

	primes := [4]int{1,2,3,4}
	fmt.Println(primes)

}