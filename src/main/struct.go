package main

import "fmt"

type Student struct {
	name string
	age int
	score float32
}
var student = new(Student)


func main() {
	student.name="zhangsan"
	student.age=22
	student.score=66.5
	//test_struct()
	pointer1 := ruturn_structPointer1()
	pointer2 := ruturn_structPointer2()
	fmt.Println(student,pointer1,pointer2)

	pointer1.name="lisi"
	fmt.Println(student,pointer1,pointer2)
}

func test_struct()  {

	var struct1 Student =  Student{"zhangsan",22,89.5}

	re := fmt.Sprintf("%T,%v\n",struct1,struct1)
	fmt.Printf(re)
	fmt.Println(struct1.name)
	fmt.Println(struct1.age)
	fmt.Println(struct1.score)

}




func ruturn_structPointer1()  *Student{

	return student
}

func ruturn_structPointer2()  *Student{

	return student
}
