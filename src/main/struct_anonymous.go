package main

import "fmt"

func main() {
	test_struct_anonymous()
	test_struct_anonymous2()
	test_struct_anonymous3()
}

func test_struct_anonymous() {

	stu := struct {
		name  string
		age   int
		score float32
	}{"zhangsan", 23, 89.9}

	fmt.Printf("%T,%v\n", stu, stu)
}

func test_struct_anonymous2() {

	stu := struct {
		string
		int
		float32
	}{"lisi", 11, 67.9}

	fmt.Printf("%T,%v\n", stu, stu)
}

func test_struct_anonymous3() {

	stu := struct {
		string
		int
		float32
	}{
		string: "zahnsan",
		int: 7,
		float32: 78.6,
	}

	fmt.Printf("%T,%v\n", stu, stu)
}
