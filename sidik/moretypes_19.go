package sidik

import "fmt"

type Vertexx struct {
	Lat, Long float64
}

var m map[string]Vertexx

func type19() {
	m = make(map[string]Vertexx)
	m["Bell Labs"] = Vertexx{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])

}