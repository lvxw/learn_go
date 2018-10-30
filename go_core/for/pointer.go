package main

import (
	"fmt"
	"reflect"
)

func main() {
	b :=12
	var a *int = &b

	fmt.Println(a == nil)

	arr := [5]int{}

	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(arr))


	parr := []*int{a,&b}

	fmt.Println(*parr[1])


	arr2 := [2]int{}

	c:=&arr2
	fmt.Println(c)
}