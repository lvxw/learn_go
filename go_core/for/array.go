package main

import "fmt"

func main()  {

	var arr = [10]int{}

	for i:=0; i<10 ; i++  {
		arr[i] = i*i*i
	}


	fmt.Println(arr)
}