package main

import "fmt"

func main() {
	result := test_function("hello world !")
	fmt.Println(result)

	a, _, b := test_function2()
	fmt.Printf("%v,%v\n",a,b)
}

func test_function(str string) (result string) {

	result = str+"*************"

	return
}

func test_function2() (int,int,int)  {

	return 1,2,3
}