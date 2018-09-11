package main

import "fmt"


func main() {
	test_type()
}

func test_type(){
	type  myInt int
	var a myInt = 10
	fmt.Printf("%T,%v\n",a,a)
	fmt.Println(a+10)
}
