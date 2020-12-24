package ds

/* TODO:
Allow duplicates
*/

// Order is an slice of sorted int (increasing order)
type Order []int

// NewOrder returns a new instance of Order
func NewOrder() *Order {
	o := make(Order, 0)
	return &o
}

// NewOrderFromInts foo
func NewOrderFromInts(n ...int) *Order {
	o := make(Order, len(n))
	copy(o, n)
	o.sort()
	return &o
}

// NewOrderFromSlice returns an order with the element in the slice
func NewOrderFromSlice(n []int) *Order {
	o := make(Order, len(n))
	copy(o, n)
	o.sort()
	return &o
}

// IsEmpty returns true if order is empty
func (o *Order) IsEmpty() bool {
	return o.Len() == 0
}

// Len returns the order's length
func (o *Order) Len() int {
	return len(*o)
}

// Get return the element at the given position
func (o *Order) Get(i int) int {
	return (*o)[i]
}

// Min return the first element
func (o *Order) Min() int {
	return (*o)[0]
}

// Max returns the last elenet
func (o *Order) Max() int {
	return (*o)[len(*o)-1]
}

//IsValid returns true if Order has all elements sorted in increasing order
func (o *Order) IsValid() bool {
	if len(*o) == 0 {
		return true
	}
	n := (*o)[0]
	for i := 1; i < len(*o); i++ {
		if (*o)[i] < n {
			return false
		}
		n = (*o)[i]
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
		pos = (pos * -1) - 1
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
	// return append(slice[:s], slice[s+1:]...)
	if pos := (*o).Search(x); pos >= 0 {
		*o = append((*o)[:pos], (*o)[pos+1:]...)
	}
}

// Search returns the position of the given element
// if the element is not present it returns a negative number:
// -(potential position)-1
func (o *Order) Search(x int) int {
	return o.binarySearch(0, len(*o)-1, x)
}

func (o *Order) binarySearch(l, r, x int) int {
	mid := (r-l)/2 + l
	if r >= l {
		if (*o)[mid] == x {
			return mid
		}
		if (*o)[mid] > x {
			return o.binarySearch(l, mid-1, x)
		}
		return o.binarySearch(mid+1, r, x)
	}
	return -mid - 1
}

func (o *Order) sort() {
	o.quick3(0, len(*o)-1)
}

func (o *Order) quick3(lo, hi int) {
	if hi <= lo {
		return
	}
	lt, i, gt := lo, lo+1, hi
	p := (*o)[lo]
	for i <= gt {
		cmp := (*o)[i]
		if cmp < p {
			o.swap(lt, i)
			lt++
			i++
		} else if cmp > 0 {
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
	(*o)[i], (*o)[j] = (*o)[j], (*o)[i]
}
