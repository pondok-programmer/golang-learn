package sidik

import "fmt"

type Vertex4 struct{
	Lat, Long float64
}

var mn = map[string]Vertex4{
	"Bell Labs" : Vertex4{
		40.68433, -74.39967,
	},
	"Google Inc" : Vertex4{
		37.42202, -122.08408,
	},
}

func type20() {
	fmt.Println(mn)
}