package main

import (
	"08-Struct/Computer"
	"fmt"
)

type Address struct {
	city, state string
}
type Person struct {
	name string
	age  int
	Address
}

func (p Person) changeName(name string) {
	//value argument methon canó't chage struct member
	p.name = name
}
func (p *Person) changeAge(age int) {
	p.age = age
}
func main() {
	var c Computer.Spec
	c.Maker = "124"
	fmt.Println(c)

	p := Person{
		Address: Address{
			city:  "shanghai",
			state: "work",
		},
		name: "jared",
		age:  18,
	}
	fmt.Println(p)

	p.changeName("tao")
	fmt.Println("after change name ", p)
	p.changeAge(123)
	fmt.Println("after chage age", p)
}
