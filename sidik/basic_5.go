package sidik

import "fmt"

func plus(x, y int) int { //jika return values hanya satu maka penulisan type return cukup satu x+y terhitung satu, jika 2 maka pemisahnya adalah koma(,)
	return x+y
}

func state()  {
	fmt.Println(plus(50, 12))
}