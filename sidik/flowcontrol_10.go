package sidik

import (
	"fmt"
	"time"
)

func flow10()  {
	fmt.Println("When is Saturday")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today+1:
		fmt.Println("Tomorrow")
	case today+2:
		fmt.Println("in Two Days")
	default:
		fmt.Println("Too far Away")
	}
}