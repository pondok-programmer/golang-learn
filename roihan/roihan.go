package roihan

import "fmt"
import "math"
import "math/rand"

func Main()  {
	fmt.Println("Hello World from Roihan")

	basics1()
	basics2()
}


func basics1() {
	fmt.Println("My favorite number is" , rand.Intn(10))
}

func basics2()  {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}