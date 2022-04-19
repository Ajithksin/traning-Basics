package main

import "fmt"

type all interface {
	baby_boy()
	baby_girl()
}

type family struct {
	father string
	mother string
}

func (f family) all() string {
	return "this " + f.father + " and " + f.mother + " has a family" + f.baby_boy() + f.baby_girl()

}

func (f family) baby_boy() string {
	return f.father + f.mother + " has a baby boy  name  don number 1"

}
func (f family) baby_girl() string {
	return "their girl is a well known doctor  ,looks likes  her  mother " + f.mother
}

func main() {
	m := family{father: "ashiwin", mother: "ankita"}
	fmt.Println(m.all())
	y := family{father: "yoyo", mother: "wowo"}
	fmt.Println(y.baby_boy())
	fmt.Println(y.baby_girl())

}
