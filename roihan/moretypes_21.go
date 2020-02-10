package roihan

import "fmt"

type literals struct {
	Lat, Long float64
}

var mp = map[string]literals{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func moretype21() {
	fmt.Println(mp)
}
