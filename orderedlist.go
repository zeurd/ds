package ds

import "fmt"

//OrderedList is a list that keeps its element in order
type OrderedList struct {
	order *Order
	e     []interface{}
	compO func(interface{}) int
}

// NewOrderedList returns a new ordered List
// compO is the function to compute the order value of the elements in the list
func NewOrderedList(compO func(interface{}) int) *OrderedList {
	return &OrderedList{
		NewOrder(),
		make([]interface{}, 0),
		compO,
	}
}

// Len foo
func (o *OrderedList) Len() int {
	return len(o.e)
}

// Slice return the ordered list as a slice
func (o *OrderedList) Slice() []interface{} {
	return o.e
}

func (o *OrderedList) String() string {
	return fmt.Sprintf("%v", o.e)
}

// Add adds x to the ordered list and returns the position it was inserted at
func (o *OrderedList) Add(x interface{}) int {
	pos := o.order.Add(o.compO(x))
	o.e = append(o.e, x)
	copy(o.e[pos+1:], o.e[pos:])
	o.e[pos] = x
	return pos
}

// AddAll adds x to the ordered list and returns the position it was inserted at
func (o *OrderedList) AddAll(xs ...interface{}) {
	for _, x := range xs {
		pos := o.order.Add(o.compO(x))
		o.e = append(o.e, x)
		copy(o.e[pos+1:], o.e[pos:])
		o.e[pos] = x
	}
}

// Delete x if is is the ordered list
func (o *OrderedList) Delete(x interface{}) int {
	if pos := (o.order).Search(o.compO(x)); pos >= 0 {
		//*o.order = append((*o.order)[:pos], (*o.order)[pos+1:]...)
		o.order.Delete(pos)
		o.e = append(o.e[:pos], o.e[pos+1:]...)
		return pos
	}
	return -1
}

// Search returns the position ox x in the list or if, absent, - potential positon - 1
func (o *OrderedList) Search(x interface{}) int {
	return o.order.Search(o.compO(x))
}

// Get return the element at the given position
func (o *OrderedList) Get(i int) interface{} {
	return o.e[i]
}

// Min return the first element
func (o *OrderedList) Min() interface{} {
	return o.e[0]
}

// Max returns the last elenet
func (o *OrderedList) Max() interface{} {
	return o.e[len(o.e)-1]
}

//IsValid returns true if Order has all elements sorted in increasing order
func (o *OrderedList) IsValid() bool {
	return o.order.IsValid()
}
