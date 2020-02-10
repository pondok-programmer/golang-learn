package roihan

import "fmt"

type Verte struct {
	Lat, Long float64
}

var ma = map[string]Verte{
	"Bell Labs": Verte{
		40.648433, -74.39967,
	},
	"Google": Verte{
		37.42202, -122.08408,
	},
}

func moretype20() {
	fmt.Println(ma)
}
