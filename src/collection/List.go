package collection

type Lists interface {
	Add(elements ...interface{})
	Insert(index int,elements ...interface{})
	Remove(index int)
	Update(index int, ele interface{})
	Get() interface{}
	GetByIndex(index int) (interface{},bool)
	Size() int
}