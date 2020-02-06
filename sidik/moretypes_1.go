package sidik

import "fmt"

func type1()  {
	i := 1000
	p := &i
	fmt.Println("i before = ",i)

	*p = 2000

	fmt.Println("i after = ",i)

}