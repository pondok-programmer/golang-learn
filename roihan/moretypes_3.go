package roihan

import "fmt"

type moretype struct {
	X int
	Y int
}

func moretype3() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
