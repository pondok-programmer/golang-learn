package roihan

import "fmt"

type vertex struct {
	Lat, Long float64
}

var m map[string]vertex

func moretype19() {
	m = make(map[string]vertex)
	m["Bell Labs"] = vertex{
		40.648433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
