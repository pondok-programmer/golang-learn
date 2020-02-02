package sidik

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 /9
	y = sum - x
	return
}

func basic7()  {
	fmt.Println(split(20))
}