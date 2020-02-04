package roihan

import (
	"fmt"
	"runtime"
)

func flowcontrol9() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Printf("%s. \n", os)
	}
}
