package sidik

import "fmt"

type Vertex5 struct {
	Lat, Long float64
}

var maa = map[string]Vertex5{
	"Bell Labs": {40.68433, -74.39967},
	"Google Inc" : {37.42202, -122.08408},
}

func type21() {
	fmt.Println(maa)
}