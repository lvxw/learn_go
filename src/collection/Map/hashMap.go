package Map

type Element struct {
	key interface{}
	value interface{}
}

type HashMap struct {
	slice []*[]Element
	size int
}