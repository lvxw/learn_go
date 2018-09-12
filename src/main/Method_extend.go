package main

import "fmt"

type Anmial struct {
	name string
	color string
}
func (a Anmial)printInfo()  {
	fmt.Printf("%v,%v\n",a.name,a.color)
}

type Dog struct {
	Anmial
	age int
}
func (d Dog)printInfo()  {
	fmt.Printf("%v,%v,%v\n",d.name,d.color,d.age)
}



func main() {
	dog := Dog{
		Anmial:Anmial{"lele", "white"},
		age: 10,
	}

	dog.printInfo()
}