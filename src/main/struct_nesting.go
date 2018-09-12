package main

import "fmt"

type Address struct {
province string
city string
}

type Person struct {
	name string
	age int
	address Address
}

func main() {
	person := Person{"zahngsan", 22, Address{"hubei", "wuhan"}}
	fmt.Printf("%T,%v\n",person,person)
	fmt.Printf("%v\n",person.name)
	fmt.Printf("%v\n",person.age)
	fmt.Printf("%v\n",person.address.province)
	fmt.Printf("%v\n",person.address.city)
}