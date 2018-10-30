package main

import "fmt"

type ArrayList struct {
	arr []int
	size int
}

func (a *ArrayList) add(el ...int){
	for i:=0;i<len(el);i++{
		a.arr = append(a.arr, el[i])
		a.size += 1
	}
}

func main() {
	list := new(ArrayList)

	list.add(1,2,3,4,5,7,8,9,10)

	fmt.Println(list.arr)
	fmt.Println(list.size)
}
