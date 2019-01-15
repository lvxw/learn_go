package lists

import (
  "errors"
  "fmt"
)

type Queue struct {
    data []interface{}
    head int
    tail int
    cap int
    size int
}

func NewQueue(cap int) Queue{
  data := make([]interface{},cap)
  return Queue{data, 0, 0,cap,0}
}

func (q *Queue) EnQueue(ele interface{}) error{
  if (q.tail+1)%q.cap == q.head{
    return errors.New("队列已满，入队失败")
  }else{
    q.data[q.tail] = ele
    q.tail = (q.tail+1)%q.cap
    q.size ++
    return nil
  }
}

func (q *Queue) DeQueue() (interface{},error) {
  if q.head == q.tail{
    return nil,errors.New("队列已空，出队失败")
  }else{
    e := q.data[q.head]
    q.data[q.head]=nil
    q.head = (q.head+1)%q.cap
    q.size --

    return e,nil
  }
}

func (q *Queue) Clear() {
  q.data = make([]interface{},q.cap)
  q.head = 0
  q.tail = 0
  q.size = 0
}

func (q *Queue) IsEmpty() bool{
  if q.head == q.tail{
    return true
  }else{
    return false
  }
}
func (q *Queue) IsFull() bool{
  if (q.tail+1)%q.cap == q.head{
    return true
  }else{
    return false
  }
}

func (q *Queue) Println(){
  fmt.Println(q.data,q.size)
}


