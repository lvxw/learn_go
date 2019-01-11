package lists

import (
	"errors"
	"fmt"
)

type Element struct {
	pre *Element
	data interface{}
	next *Element
}

type LinkedList struct {
	root Element
	size int
}

func NewLinkedList() *LinkedList{
	root := Element{nil,nil,nil}
	linkedList := LinkedList{root,0}
	return &linkedList
}

func (linkedList *LinkedList) Println(){
	s := make([]interface{},0)

	element := linkedList.root.next

	for element != nil{
		s = append(s, element.data)
		element = element.next
	}

	fmt.Println(s,linkedList.size)
}

func (linkedList *LinkedList) getLastNodePointer() *Element {
	tmpNode := &(linkedList.root)
	for {
		re := tmpNode.next
		if re == nil{
			break
		}
		tmpNode = re
	}
	return tmpNode
}

func (linkedList *LinkedList) getNodeByIndex(index int) *Element {
	tmpNode := &(linkedList.root)
	for i:=0;i<=index;i++{
		tmpNode = tmpNode.next
	}
	return tmpNode
}

func (linkedList *LinkedList) AddHead(ele ...interface{})  {
	for i:= len(ele)-1; i>=0; i--{
		node := Element{nil,ele[i],nil}
		if linkedList.root.next == nil{
			linkedList.root.next = &node
			node.pre = &(linkedList.root)
		}else{
			tmp := linkedList.root.next
			node.next = tmp
			tmp.pre = &node
			node.pre = &(linkedList.root)
			linkedList.root.next = &node
		}
		linkedList.size += 1
	}
}

func (linkedList *LinkedList) AddTail(ele ...interface{})  {
	tmpNode := linkedList.getLastNodePointer()

	for _,e := range ele{
		node := Element{nil,e,nil}
		tmpNode.next = &node
		node.pre = tmpNode
		tmpNode = &node
		linkedList.size += 1
	}
}

func (linkedList *LinkedList) Get(index int) (interface{},error) {
	if index < 0 || index >= linkedList.size {
		return nil,errors.New("index out of range")
	}else{
		tmpNode := linkedList.getNodeByIndex(index)
		return tmpNode.data,nil
	}

}

func (linkedList *LinkedList) Set(index int, ele interface{}) error {
	if index < 0 || index >= linkedList.size {
		return errors.New("index out of range")
	}else{
		tmpNode := linkedList.getNodeByIndex(index)
		tmpNode.data = ele
		return nil
	}
}

func (linkedList *LinkedList) Remove(index int) (interface{},error) {
	if index < 0 || index >= linkedList.size {
		return nil,errors.New("index out of range")
	}else{
		tmpNode:=linkedList.getNodeByIndex(index)
		if index == linkedList.size -1{
			tmpNode.pre.next = nil
			tmpNode.pre=nil
		} else{
			tmpNode.pre.next = tmpNode.next
			tmpNode.next.pre = tmpNode.pre
			tmpNode.pre = nil
			tmpNode.next=nil
		}
		linkedList.size--
		return tmpNode.data,nil
	}
}

func (linkedList *LinkedList) InsertBefore(index int, ele ...interface{}) error {
	if index < 0 || index >= linkedList.size {
		return errors.New("index out of range")
	}else{
		if index == 0{
			linkedList.AddHead(ele...)
		}else{
			tmpNode := linkedList.getNodeByIndex(index-1)

			afterNode := tmpNode.next
			for _,e := range ele{
				node := Element{nil,e,nil}
				tmpNode.next = &node
				node.pre = tmpNode
				tmpNode = &node
				linkedList.size += 1
			}

			tmpNode.next = afterNode
			afterNode.pre = tmpNode
		}
		return nil
	}
}

func (linkedList *LinkedList) InsertAfter(index int, ele ...interface{}) error {
	if index < 0 || index >= linkedList.size {
		return errors.New("index out of range")
	}else{
		if index == linkedList.size-1{
			linkedList.AddTail(ele...)
		}else{
			tmpNode := linkedList.getNodeByIndex(index)

			afterNode := tmpNode.next
			for _,e := range ele{
				node := Element{nil,e,nil}
				tmpNode.next = &node
				node.pre = tmpNode
				tmpNode = &node
				linkedList.size += 1
			}

			tmpNode.next = afterNode
			afterNode.pre = tmpNode
		}
		return nil
	}
}
