package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
}

func main() {
	person := Person{"zhangsan", 12}
	fmt.Println(person)
	person.name="lisi"
	person.age=23

}
