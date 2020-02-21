package sidik

import "fmt"

func type9() {
	v := []int{4, 5, 6, 33, 5, 7}
	fmt.Println(v)

	p := []bool{true, false, false, true, false}
	fmt.Println(p)


	s := []struct {
		umur int
		fact bool
	}{
		{9, true},
		{11, true},
		{4, false},
		{6, true},
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
