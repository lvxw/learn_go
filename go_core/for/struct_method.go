package main

import "fmt"

type Animal struct{
	name string
	color string
}
func (a *Animal) construct(name string,color string){
	a.name = name
	a.color = color
}

func main() {
	animal := new(Animal)
	animal.construct("panda","white-black")

	fmt.Println(*animal)
}
