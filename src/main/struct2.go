package main

import "fmt"

type Student2 struct {
	name string
	age int
	score float32
}

func main() {

	stu1 := Student2{"zhangsan", 23, 99.5}

	stu2 := Student2{
		"zhangsan",
		23,
		99.5,
	}

	stu3 := Student2{
		age:23,
		name:"zhangsan",
		score:99.5,
	}

	stu4 := new(Student2)

	fmt.Printf("%T,%v\n",stu1,stu1)
	fmt.Printf("%T,%v\n",stu2,stu2)
	fmt.Printf("%T,%v\n",stu3,stu3)
	fmt.Printf("%T,%v\n",stu4,stu4)
}
