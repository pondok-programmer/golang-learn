package sidik

import (
	"fmt"
	"math"
)

func pow(x, n , lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}

	return lim
}

func flow6()  {
	fmt.Println(
		pow(9, 4, 10),
		pow(3, 4, 10),
	)			
}