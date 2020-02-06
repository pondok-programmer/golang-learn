package roihan

import (
	"fmt"
	"time"
)

func flowcontrol11() {
	s := "Roihan"
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!", s)
	case t.Hour() < 17:
		fmt.Println("Good Aftrenoon", s)
	default:
		fmt.Println("Good Evening", s)
	}
}
