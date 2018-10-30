package main

import "fmt"

// 匿名结构体
var person = struct {
	name string
	age int
}{"lisi",45}

// 匿名结构体\匿名字段
var person2 = struct {
	string
	int
}{"wangwu",23}

func main() {
	fmt.Println(person.name)
	fmt.Println(person.age)

	fmt.Println(person2.string)
	fmt.Println(person2.int)
}
