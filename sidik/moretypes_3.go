package sidik

import "fmt"

type Vertex1 struct {
	X int
	Y int
	Z int
}

func type3() { //mengakses properti struct dengan titik
	p := Vertex{8, 4, 10}
	p.X = 19
	p.Z	= 8

	fmt.Println(p.X, p.Z)

}