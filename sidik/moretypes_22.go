package sidik

import "fmt"

func type22() {
	mm := make(map[string]int)

	mm["Answer"] = 42
	fmt.Println("The Value: ", mm["Answer"])

	mm["Answer"] = 48
	fmt.Println("The Value: ", mm["Answer"])

	delete(mm, "Answer")
	fmt.Println("The Value: ", mm["Answer"])

	v, ok := mm["Answer"]
	fmt.Println("The Value: ", v ,"Present", ok)
}