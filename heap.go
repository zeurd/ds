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
	Value(interface{}) (int, bool)
	String() string
	Subtree(int) [][]interface{}
	IsValid()
	//Slice() []interface{}
}

// NewHeap returns a new heap
func NewHeap() Heap {
	return newHeap()
}

// NewHeapWithEvalFunction returns a new heap that uses the given function to evaluate priority in the heap
func NewHeapWithEvalFunction(f func(x interface{})int ) Heap {
	return newHeapWithEval(f)
}
