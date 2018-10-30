package main

import "fmt"

type Student struct{
	school Schhool
	gread string
	name string
}

type Schhool struct{
	address string
	name string
}



func main() {
	student := new(Student)

	student.school=Schhool{"海淀区","北大附小"}
	student.gread="six"
	student.name="张三"

	fmt.Println(*student)
}
