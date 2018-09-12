package main

import (
	"container/list"
	"fmt"
)

func main()  {
	my_list()
	my_list2()
}

func my_list()  {
	li := list.New()
	li.PushBack(1)
	li.PushBack(2)
	li.PushBack("hello")

	fmt.Printf("%T,%v\n",li,li)
}

func my_list2()  {
	var li = list.List{}
	fmt.Printf("%T,%v\n",li,li)

}