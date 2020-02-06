package roihan

import "fmt"

func flowcontrol3() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("Hasilnya :", sum)
}
