package roihan

import "fmt"

type Moretype struct {
	X int
	Y int
}

func moretype4() {
	v := Moretype{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
