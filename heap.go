package ds

// Heap defines an addressable heap interface
type Heap interface {
	Pop() interface{}
	Peek() interface{}
	Push(interface{})
	Insert(interface{}, int)
	Update(interface{}, int)
	Delete(interface{})
	IsEmpty() bool
	Key(interface{}) (int, bool)
	String() string
	Subtree(int) [][]interface{}
	IsValid()
	//Slice() []interface{}
}

// NewHeap returns a new heap
func NewHeap(duplicate bool) Heap {
	return newHeap(duplicate)
}

// NewHeapWithEvalFunction returns a new heap that uses the given function to evaluate priority in the heap
func NewHeapWithEvalFunction(duplicate bool, f func(x interface{})int) Heap {
	return newHeapWithEval(duplicate, f)
}
