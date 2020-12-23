package ds

import (
	"math/rand"
)

// Order is an slice of sorted int (increasing order)
type Order struct {
	o []int
}

// NewOrder returns a new instance of Order
func NewOrder() *Order {
	return &Order{make([]int, 0)}
}

// NewOrderFromSlice returns an order with the element in the slice
func NewOrderFromSlice(n []int) *Order {
	o := &Order{n}
	o.quicksort()
	return o
}

//Add adds a new int to Order
func (o *Order) Add(x int) {
	o.o = append(o.o, x)
}

// Search returns the position of the given element or -(position where it should be)
func (o *Order) Search(x int) int {
	return o.binarySearch(0, len(o.o), x)
}

func (o *Order) binarySearch(l, r, x int) int {
	mid := l + (r-1)/2
	if r >= l {
		if o.o[mid] == x {
			return mid
		}
		if o.o[mid] > x {
			return o.binarySearch(l, mid-1, x)
		}
		return o.binarySearch(mid+1, r, x)
	}
	//special case, x is not present but should be first element
	if -mid == 1 {
		return -1
	}
	return 1
}

func (o *Order) quicksort() {
	o.qs(0, len(o.o))
}
func (o *Order) qs(l, r int) {
	if r-l <= 1 {
		return
	}
	//chose pivot element
	p := rand.Intn(r-l) + l
	//swap pivot with 1st element
	o.swap(p, l)

	//partition around pivot
	split := o.partition(l, r)
	//recursively sort 1st part
	o.qs(l, split-1)
	//recursively sort 2ndrt
	o.qs(split, r)
	//No combining part!
}

func (o *Order) partition(l, r int) int {
	//we swapped pivot with l
	pivot := o.o[l]
	i := l + 1
	for j := l + 1; j < r; j++ {
		//can add extra condition: && we've already seen element bigger than the pivot
		if o.o[j] < pivot {
			//if new element is less than the pivot,
			//swap with ith element, the left-most element less that's bigger than the pivot
			o.swap(i, j)
			i++
		} //if new element is bigger than the pivot, nothing to do
	}
	//swap the pivot in its correst position
	o.swap(l, i-1)
	return i
}

func (o *Order) swap(i, j int) {
	o.o[i], o.o[j] = o.o[j], o.o[i]
}
