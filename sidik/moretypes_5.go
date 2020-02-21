package sidik

import "fmt"

type Vertex3 struct {
	X, Y int
}

var (
	v = Vertex3{5, 44} // struct asli
	v2 = Vertex3{X:6} // mengakses satu property di struct
	v3 = Vertex3{} // struct zzero value
	p = &Vertex3{3,4} //pointer struct, harus dengan * jika ingin value yang asli 
)

func type5() {
	fmt.Println(v, v2, v3, p)
}