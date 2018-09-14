package main

import "fmt"

type LinkedList struct {
	data interface{}
	pNext *LinkedList
}

func main() {
	var linkedList = new(LinkedList)

	linkedList.set(1,2,3,4,5,6,7,8,9,10)
	linkedList.set("spark")
	linkedList.forEach()
}


func (head *LinkedList) setBefore(slice ...interface{})  {
	for i:=0;i<len(slice);i++{
		node := new(LinkedList)
		node.data=slice[i]
		if head.pNext != nil{
			tmp := head.pNext
			head.pNext = node
			node.pNext = tmp
		}else{
			head.pNext = node
		}
	}
}

func (head *LinkedList) set(slice ...interface{})  {
	for i:=0;i<len(slice);i++{
		node := new(LinkedList)
		node.data=slice[i]

		if head.pNext ==nil{
			head.pNext = node
		}else{
			chNode := head.pNext
			for chNode.pNext != nil  {
				chNode = chNode.pNext
			}
			chNode.pNext = node
		}
	}
}

func (head *LinkedList) get(index int)  interface{}{

	if index < 0{
		return nil
	}
	
	count := 0
	
	node := head.pNext
	for node != nil{
		if count == index{
			return node.data
		}else{
			count ++
			node = node.pNext
		}
	}

	return nil
}

func (head *LinkedList) forEach() {
	node := head.pNext
	for node != nil{
		fmt.Println(node.data)
		node = node.pNext
	}
}
