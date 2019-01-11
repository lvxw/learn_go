package lists

import (
	"errors"
)

type ArrayList struct {
	data []interface{}
	size int
}

func NewArrayList(ele ...interface{}) *ArrayList{

	if len(ele)==0{
		return &(ArrayList{make([]interface{},0),0})
	}else{
		s := make([]interface{},len(ele))
		for i,e := range ele{
			s[i] = e
		}
		return &(ArrayList{s,len(ele)})
	}


	return &(ArrayList{[]interface{}{},0})
}

func (arrayList *ArrayList) Add(ele ...interface{})  {
	for _,e := range ele{
		arrayList.data = append(arrayList.data, e)
		arrayList.size += 1
	}
}
func (arrayList *ArrayList) Insert(index int, ele ...interface{}) error {
	if index < 0 || index >= arrayList.size {
		return errors.New("index out of range")
	}else{
		s := make([]interface{},0)
		s = append(s, arrayList.data[:index]...)
		s = append(s, ele...)
		s = append(s, arrayList.data[index:]...)
		arrayList.data = s
		arrayList.size += len(ele)
		return nil
	}
}

func (arrayList *ArrayList) Remove(index int) (interface{}, error) {
	if index < 0 || index >= arrayList.size {
		return nil,errors.New("index out of range："+string(index))
	}else{
		removeEle := arrayList.data[index]
		arrayList.data = append(arrayList.data[:index],arrayList.data[index+1:]...)
		arrayList.size--
		return removeEle,nil
	}
}

func (arrayList *ArrayList) Clear()  {
	arrayList.data = make([]interface{},0)
	arrayList.size = 0
}

func (arrayList *ArrayList) Set(index int,ele interface{}) (interface{}, error) {

	if index < 0 || index >= arrayList.size{
		return nil,errors.New("index out of range")
	}else{
		updateEle := arrayList.data[index]
		arrayList.data[index] = ele
		return updateEle,nil
	}
}

func (arrayList *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= arrayList.size{
		return nil,errors.New("index out of range")
	}else{
		return arrayList.data[index], nil
	}
}


