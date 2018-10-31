package main

import (
	"fmt"
	"go/types"
)

type A interface {
}

type Cat3 struct {
	name string
	age string
}

type Person3 struct {
	name string
	sex string
}




func main() {

	var a1 A = Cat3{}
	var a2 A = Person3{}
	var a3 A = ""
	fmt.Println(s)

	fmt.Printf("%T\n",a1)
	fmt.Printf("%T\n",a2)
	fmt.Printf("%T\n",a3)

}
