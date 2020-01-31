package roihan

import "fmt"
import "math/rand"

func Main()  {
	fmt.Println("Hello World from Roihan")
	basics1()
	basics2()
	basics3()
	basics4()
	basics5()
}


func basics1() {
	fmt.Println("My favorite number is" , rand.Intn(10))
}
