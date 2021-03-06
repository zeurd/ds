package ds

// BinarySearchTree is an AVL BST
type BinarySearchTree interface {
	Search(int) interface{}
	Insert(int, interface{}) bool
	Push(interface{})
	Delete(int) bool
	DeleteKV(int, interface{}) bool
	Predecessor(int) interface{}
	Min() interface{}
	Max() interface{}
	MinK() (int, interface{})
	MaxK() (int, interface{})
	Len() int
	Height() int
	IsValid() bool
	Slice() []interface{}
	String() string
}

// NewBinarySearchTree returns a new BST
// duplicate = true to allow duplicate keys
// kf, the key function to evaluate the key if intend to use Push(x) and not Insert(k,x)
func NewBinarySearchTree(duplicate bool, kf func(x interface{}) int) BinarySearchTree {
	if !duplicate {
		return newBst(kf)
	}
	 return newBstDup(kf)
}
