package sidik

import (
	"fmt"
	"math/cmplx"
)

var ( //membuat variabel bisa sekaligus dengan pembungkus
	Tobe bool	= false
	Maxint uint64 = 1<<64 - 1
	zo	complex128 = cmplx.Sqrt(-5 + 12i)
)

func basic11()  {
	fmt.Printf("Type: %T values: %v\n", Tobe, Tobe)
	fmt.Printf("Type: %T values: %v\n", Maxint, Maxint)
	fmt.Printf("Type: %T values: %v\n", zo, zo)
}