package ds

// BstD is a BST that can contain duplicate key
type BstD struct {
	d map[int]int //duplicates
	b *Bst
}

//NewBstD foo
func NewBstD() *BstD {
	return &BstD{
		make(map[int]int),
		NewBst(),
	}
}

//NewBstDWithKeyFunc foo
func NewBstDWithKeyFunc(kf func(interface{})int) *BstD{
	return &BstD{
		make(map[int]int),
		NewBstWithKeyFunc(kf),
	}
}

func Push()
