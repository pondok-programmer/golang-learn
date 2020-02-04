package roihan

import (
	"fmt"
	"math"
)

func pow_control(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func flowcontrol7() {
	fmt.Println("Hasilnya :",
		pow(3, 3, 10),
		pow(3, 3, 20),
	)
}
