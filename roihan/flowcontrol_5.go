package roihan

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func flowcontrol5() {
	fmt.Println("Hasilnya :", sqrt(2), sqrt(-4))
}
