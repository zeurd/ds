package ds

import (
	"math/rand"
)

// Order is an slice of sorted int (increasing order)
type Order []int

// NewOrder returns a new instance of Order
func NewOrder() Order {
	return make([]int, 0)
}

// NewOrderFromSlice returns an order with the element in the slice
func NewOrderFromSlice(n []int) Order {
	var o Order
	copy(o, n)
	o.quicksort()
	return o
}

//IsValid returns true if Order has all elements sorted in increasing order
func (o Order) IsValid() bool {
	if len(o) == 0 {
		return true
	}
	n := o[0]
	for i := 1; i < len(o); i++ {
		if o[i] < n {
			return false
		}
		n = o[i]
	}
	return true
}

//Add adds a new int to Order
func (o *Order) Add(x int) {
	if len(*o) == 0 {
		*o = []int{x}
		return
	}
	if pos := o.Search(x); pos < 0 {
		pos = (pos*-1)-1
		*o = append(*o, x)
		copy((*o)[pos+1:], (*o)[pos:])
		(*o)[pos] = x
	}
}

// Delete deletes x if it is in the order
func (o *Order) Delete(x int) {
	if len(*o) == 0 {
		return
	}
	if pos := o.Search(x); pos >= 0 {
		end := len(*o) - 1
		o.swap(pos, end)
		*o = (*o)[:end]
	}
}


// Search returns the position of the given element
// if the element is not present it returns a negative number:
// -(potential position)-1
func (o Order) Search(x int) int {
	return o.binarySearch(0, len(o)-1, x)
}

func (o Order) binarySearch(l, r, x int) int {
	mid := (r-l)/2 + l
	//fmt.Printf("l: %d. r: %d. mid: %d\tfrom: %v\n", l, r, mid, f)
	if r >= l {
		if o[mid] == x {
			return mid
		}
		if o[mid] > x {
			return o.binarySearch(l, mid-1, x)
		}
		return o.binarySearch(mid+1, r, x)
	}
	//need to do -1 for case if x would be first element
	return -mid-1
}

func (o Order) quicksort() {
	o.qs(0, len(o))
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

func (o Order) partition(l, r int) int {
	//we swapped pivot with l
	pivot := o[l]
	i := l + 1
	for j := l + 1; j < r; j++ {
		//can add extra condition: && we've already seen element bigger than the pivot
		if o[j] < pivot {
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

func (o Order) swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}
