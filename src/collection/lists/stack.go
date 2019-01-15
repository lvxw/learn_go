package lists

import (
	"errors"
	"fmt"
)

type Stack struct {
	data []interface{}
	head int
	cap int
}

func NewStack(cap int) Stack{
	data := make([]interface{},cap)
	return Stack{data,0,cap}
}

func (s *Stack) Push(ele interface{}) error{
	if s.head == s.cap{
		return errors.New("栈列已满，压栈失败")
	}else{
		s.data[s.head] = ele
		s.head ++
		return nil
	}
}

func (s *Stack) Pop() (interface{},error){
	if s.head == 0{
		return nil,errors.New("栈列已空，弹栈失败")
	}else{
		e := s.data[s.head-1]
		s.data[s.head-1] = nil
		s.head --
		return e,nil
	}
}

func (s *Stack) Clear(){
	s.head = 0
	s.data = make([]interface{},s.cap)
}

func (s *Stack) IsEmpty() bool{
	return s.head == 0
}

func (s *Stack) IsFull() bool{
	return s.head+1 == s.cap
}


func (s *Stack) Size() int{
	return s.head
}

func (s *Stack) Println(){
	fmt.Println(s.data, s.head)
}
