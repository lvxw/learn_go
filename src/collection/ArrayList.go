package collection

type ArrayList struct {
	arr []interface{}
	len int
}

func (arrayList *ArrayList) Add(elements ...interface{}) ()  {
	for _,ele := range elements{
		arrayList.arr = append(arrayList.arr, ele)
		arrayList.len += 1
	}
}
func (arrayList *ArrayList) Insert(index int,elements ...interface{}) ()  {
	if index < 0{
		arrayList.arr = append(elements, arrayList.arr...)
		arrayList.len+= len(elements)
		return
	}
	if index > arrayList.len-1{
		arrayList.arr = append(arrayList.arr, elements...)
		arrayList.len+= len(elements)
		return
	}
	arrHead := arrayList.arr[0:index]
	arrTail := arrayList.arr[index+1:]

	arrayList.arr = append(arrHead, elements...)
	arrayList.arr = append(arrayList.arr, arrTail...)
	arrayList.len+= len(elements)
}

func (arrayList *ArrayList) Remove(index int){
	if index >= arrayList.len || index <0 {
		return
	}
	arrHead := arrayList.arr[0:index]
	arrTail := arrayList.arr[index+1:]

	arrayList.arr = append(arrHead, arrTail...)
	arrayList.len -=1

}
func (arrayList *ArrayList) Update(index int,ele interface{}){
	if index >= arrayList.len {
		arrayList.arr = append(arrayList.arr, ele)
		arrayList.len += 1
		return
	}
	if index < 0{
		arrayList.arr = append([]interface{}{ele},arrayList.arr...)
		arrayList.len += 1
		return
	}
	arrayList.arr[index] = ele
}
func (arrayList *ArrayList) Get() interface{}{
	return arrayList.arr[arrayList.len-1]
}
func (arrayList *ArrayList) GetByIndex(index int) (interface{},bool){
	if index >= arrayList.len || index < 0 {
		return nil,false
	}
	return arrayList.arr[index],true
}
func (arrayList *ArrayList) Size() int{
	return arrayList.len
}

