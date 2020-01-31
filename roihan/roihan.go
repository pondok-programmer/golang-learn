package roihan


import "fmt"
import "math"
import "math/rand"

func Main()  {
	fmt.Println("Hello World from Roihan")
	basic1()
	basics1()
}

func basic1()  {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

func basics1() {
	fmt.Println("My favorite number is" , rand.Intn(10))
}