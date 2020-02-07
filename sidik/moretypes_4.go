package sidik

import "fmt"

type Vertex2 struct {
	X int
	Y int
}

func type4() {
	i := Vertex2{2, 3}
	p := &i
	p.Y = 8

	fmt.Println(i)

}