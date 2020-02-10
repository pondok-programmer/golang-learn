package roihan

import "fmt"

type MoretypeVertex struct {
	X, Y int
}

var (
	v1 = MoretypeVertex{1, 2}  //Has type vertex
	v2 = MoretypeVertex{X: 1}  //Y:0 is implicit
	v3 = MoretypeVertex{}      //X:0 and Y:0
	p  = &MoretypeVertex{1, 2} //Has type *Vertex
)

func moretype5() {
	fmt.Println(v1, p, v2, v3)
}
