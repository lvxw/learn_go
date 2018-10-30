package main

import "fmt"

type Person2 struct{
	name string
	age int
}

type Student2 struct{
	Person2
	score float32
}



func main() {
	student := new(Student2)

	student.Person2 =Person2{"lisi",12}
	student.score=98.5

	fmt.Println(student)

	student2 := Student2{Person2{"we",23}, 78.5}
	fmt.Printf("%+v\n",student2)
	fmt.Println(student2.name)
	fmt.Println(student2.age)
	fmt.Println(student2.score)
}
