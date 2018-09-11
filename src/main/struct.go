package main

import "fmt"

type Student struct {
	name string
	age int
	score float32
}


func main() {
	test_struct()
}

func test_struct()  {

	var struct1 Student =  Student{"zhangsan",22,89.5}

	re := fmt.Sprintf("%T,%v\n",struct1,struct1)
	fmt.Printf(re)
	fmt.Println(struct1.name)
	fmt.Println(struct1.age)
	fmt.Println(struct1.score)

}
