package sidik

import (
	"fmt"
	"math"
)

func pow1(x, n , lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}

func flow7()  {
	fmt.Println(
		pow1(9, 4, 10),
		pow1(3, 4, 10),
	)			
}