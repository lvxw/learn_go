package main

import "fmt"

func main() {
	arr := []string{"Hadoop","Spark","Storm"}

	test_variable_param("java","scala","python","go","nodeJs")
	test_variable_param(arr...)
}

func test_variable_param(str ...string)  {
	for _,value := range str{
		fmt.Println(value)
	}
	fmt.Println("-------------------------")
}