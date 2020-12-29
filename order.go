package ds

import "sort"

// Order is an slice of sorted int (increasing order)
// /!\ in O(n) !!!
type Order struct {
	o []int
}

// NewOrder returns a new instance of Order
func NewOrder() *Order {
	o := make([]int, 0)
	return &Order{o}
}

// NewOrderFromInts foo
func NewOrderFromInts(n ...int) *Order {
	sort.Ints(n)
	return &Order{n}
}

// NewOrderFromSlice returns an order with the element in the slice
func NewOrderFromSlice(n []int) *Order {
	sort.Ints(n)
	return &Order{n}
}

// IsEmpty returns true if order is empty
func (o *Order) IsEmpty() bool {
	return o.Len() == 0
}

// Len returns the order's length
func (o *Order) Len() int {
	return len(o.o)
}

// Get return the element at the given position
func (o *Order) Get(i int) int {
	return o.o[i]
}

// Min return the first element
func (o *Order) Min() int {
	return o.o[0]
}

// Max returns the last elenet
func (o *Order) Max() int {
	return o.o[len(o.o)-1]
}

//IsValid returns true if Order has all elements sorted in increasing order
func (o *Order) IsValid() bool {
	if len(o.o) == 0 {
		return true
	}
	n := o.o[0]
	for i := 1; i < len(o.o); i++ {
		if o.o[i] < n {
			return false
		}
		n = o.o[i]
	}
	return true
}

//Add adds a new int to Order and returs the position it was inserted at
func (o *Order) Add(x int) int {
	if len(o.o) == 0 {
		o.o = []int{x}
		return 0
	}
	pos := o.Search(x)
	if pos < 0 {
		pos = (pos * -1) - 1
	}
	o.o = append(o.o, 0)
	copy(o.o[pos+1:], o.o[pos:])
	o.o[pos] = x
	// o.insert(pos, x)
	return pos
}

// Delete deletes x if it is in the order and returns the position of the delete element
func (o *Order) Delete(x int) int {
	if len(o.o) == 0 {
		return -1
	}
	if pos := o.Search(x); pos >= 0 {
		o.o = append(o.o[:pos], o.o[pos+1:]...)
		return pos
	}
	return -1
}

// Search returns the position of the given element
// if the element is not present it returns a negative number:
// -(potential position)-1
func (o *Order) Search(x int) int {
	return o.binarySearch(0, len(o.o)-1, x)
}

func (o *Order) binarySearch(l, r, x int) int {
	mid := (r-l)/2 + l
	if r >= l {
		if o.o[mid] == x {
			return mid
		}
		if o.o[mid] > x {
			return o.binarySearch(l, mid-1, x)
		}
		return o.binarySearch(mid+1, r, x)
	}
	return -mid - 1
}

func (o *Order) sort() {
	o.quick3(0, len(o.o)-1)
}

func (o *Order) quick3(lo, hi int) {
	if hi <= lo {
		return
	}
	lt, i, gt := lo, lo+1, hi
	p := o.o[lo]
	for i <= gt {
		cmp := o.o[i]
		if cmp < p {
			o.swap(lt, i)
			lt++
			i++
		} else if cmp > p {
			o.swap(i, gt)
			gt--
		} else {
			i++
		}
	}
	o.quick3(lo, lt-1)
	o.quick3(gt+1, hi)
}

func (o *Order) swap(i, j int) {
	o.o[i], o.o[j] = o.o[j], o.o[i]
}
