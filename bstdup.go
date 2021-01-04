package ds

// BstD is a BST that can contain duplicate key
type BstD struct {
	d map[int]int //duplicates
	b *BinarySearchTree
}

//NewBstD foo
