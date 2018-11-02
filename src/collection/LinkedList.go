package collection

import "fmt"

type LinkedList struct {
	previous *LinkedList
	data interface{}
	next *LinkedList
	tail *LinkedList
	len int
}

func (linkedList *LinkedList) Add(elements ...interface{}) ()  {

	for _,ele := range elements{
		list := new(LinkedList)
		list.data = ele
		list.next = nil

		tmpNode := linkedList.next
		if tmpNode == nil{
			linkedList.next = list
			list.previous = linkedList
		}else{
			for tmpNode.next !=nil{
				tmpNode = tmpNode.next
			}
			tmpNode.next = list
			list.previous = tmpNode
		}
		linkedList.tail=list

		linkedList.len += 1
	}

}

func (linkedList *LinkedList) Insert(index int, elements ...interface{}) ()  {

	tmpHeadNode := new(LinkedList)
	tmpNode := tmpHeadNode

	for _,ele := range elements{
		node := new(LinkedList)
		node.data = ele
		node.next = nil

		for tmpNode.next !=nil{
			tmpNode = tmpNode.next
		}

		tmpNode.next = node
		node.previous = tmpNode
	}

	if index < 0{
		if linkedList.tail == nil{
			linkedList.next = tmpHeadNode.next
			tmpHeadNode.next.previous = linkedList
			linkedList.tail = tmpNode.next
			linkedList.len += len(elements)
		}else{
			afterNodes := linkedList.next
			linkedList.next = tmpHeadNode.next
			tmpHeadNode.previous = linkedList
			tmpNode.next.next = afterNodes
			afterNodes.previous = tmpNode.next
			linkedList.len += len(elements)
		}
		return
	}

	if index > linkedList.len-1{
		if linkedList.tail == nil{
			linkedList.next = tmpHeadNode.next
			tmpHeadNode.next.previous = linkedList
		}else{
			linkedList.tail.next = tmpHeadNode.next
			tmpHeadNode.next.previous = linkedList.tail
		}
		linkedList.tail = tmpNode.next
		linkedList.len += len(elements)
		return
	}

	isFromHead := linkedList.len/2 >= index
	FindNode := linkedList

	if isFromHead {
		for i:=0;i<=index;i++{
			FindNode = FindNode.next
		}
	}else{
		FindNode = linkedList.tail
		for i:=linkedList.len-1;i>index;i--{
			FindNode = FindNode.previous
		}
	}

	FindNode.previous.next = tmpHeadNode.next
	tmpHeadNode.next.previous = FindNode.previous
	tmpNode.next.next = FindNode
	FindNode.previous = tmpNode.next
	linkedList.len += len(elements)

}

func (linkedList *LinkedList) AddTail(elements ...interface{}) ()  {
	linkedList.Add(elements...)
}

func (linkedList *LinkedList) AddHead(elements ...interface{}) ()  {

	for _,ele := range elements{
		tmpNode := linkedList.next

		list := new(LinkedList)
		list.data = ele

		linkedList.next =  list
		list.next = tmpNode
		if tmpNode != nil {
			tmpNode.previous = list
		}else{
			linkedList.tail = list
		}
		list.previous = linkedList

		linkedList.len +=1
	}
}

func (linkedList *LinkedList) Remove(index int){
	if index > linkedList.len-1 || index < 0{
		return
	}

	if index == 0{
		tmpNode := linkedList.next
		linkedList.next = tmpNode.next
		tmpNode.next.previous = linkedList
		return
	}

	isFromHead := linkedList.len/2 >= index
	tmpNode := linkedList

	if isFromHead {
		for i:=0;i<=index;i++{
			tmpNode = tmpNode.next
		}
		tmpNode.previous.next = tmpNode.next
		tmpNode.next.previous =  tmpNode.previous
	}else{
		tmpNode = linkedList.tail
		for i:=linkedList.len-1;i>index;i--{
			tmpNode = tmpNode.previous
		}
		if index == linkedList.len-1{
			tmpNode.previous.next = nil
			linkedList.tail = tmpNode.previous
			tmpNode.previous = nil
		}else {
			tmpNode.previous.next = tmpNode.next
			tmpNode.next.previous =  tmpNode.previous
		}
	}
	linkedList.len -= 1
}

func (linkedList *LinkedList) Update(index int,ele interface{}){
	if index > linkedList.len-1{
		node := new(LinkedList)
		node.next = nil
		node.data = ele

		linkedList.tail.next = node
		node.previous = linkedList.tail
		linkedList.tail = node

		linkedList.len+=1
		return
	}
	if index < 0{
		node := new(LinkedList)
		node.data = ele

		tmpNode := linkedList.next
		linkedList.next = node
		node.next = tmpNode
		node.previous = linkedList

		linkedList.len+=1
		return
	}
	isFromHead := linkedList.len/2 >= index
	tmpNode := linkedList

	if isFromHead {
		for i:=0;i<=index;i++{
			tmpNode = tmpNode.next
		}
	}else{
		tmpNode = linkedList.tail
		for i:=linkedList.len-1;i>index;i--{
			tmpNode = tmpNode.previous
		}
	}
	tmpNode.data = ele
}

func (linkedList *LinkedList) Get() interface{}{
	return linkedList.next.data
}

func (linkedList *LinkedList) GetByIndex(index int) (interface{},bool){
	if index<0 || index > linkedList.len-1{
		return nil,false
	}else {
		isFromHead := linkedList.len/2 >= index
		tmpNode := linkedList

		if isFromHead {
			for i:=0;i<=index;i++{
				tmpNode = tmpNode.next
			}
		}else{
			tmpNode = linkedList.tail
			for i:=linkedList.len-1;i>index;i--{
				tmpNode = tmpNode.previous
			}
		}
		return tmpNode.data,true
	}
}

func (linkedList *LinkedList) Size() int{
	return linkedList.len
}

func (linkedList *LinkedList) ForEach(){
	tmpNode := linkedList.next
	fmt.Print("[")
	for tmpNode !=nil {
		if tmpNode.next != nil {
			fmt.Print(tmpNode.data," ")
		}else{
			fmt.Print(tmpNode.data)
		}
		tmpNode= tmpNode.next
	}
	fmt.Print("]\n")
}
