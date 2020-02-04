package roihan

import "fmt"

func flowcontrol1() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Hasilnya :", sum)
}
