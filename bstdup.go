package ds

// bstD is a BST that can contain duplicate key
type bstDup struct {
	dup map[int]Set //duplicates
	bst *bst
	len int
}

//newBstDup
func newBstDup(kf func(interface{}) int) *bstDup {
	return &bstDup{
		make(map[int]Set),
		newBst(kf),
		0,
	}
}

func (b *bstDup) Search(key int) interface{} {
	return b.bst.Search(key)
}

func (b *bstDup) Insert(key int, x interface{}) bool {
	b.len++
	if _, ok := b.dup[key]; ok {
		b.dup[key].Add(x)
		return true
	}
	b.dup[key] = NewSet()
	b.dup[key].Add(x)
	return b.bst.Insert(key, x)
}

func (b *bstDup) Push(x interface{}) {
	k := b.bst.kf(x)
	b.Insert(k, x)
}

func (b *bstDup) Delete(key int) bool {
	d, ok := b.dup[key]
	if !ok {
		return false
	}
	b.len--
	d.Pop()
	if len(d) > 0 {
		return true
	}
	return b.bst.Delete(key)
}

func (b *bstDup) DeleteKV(key int, x interface{}) bool {
	d, ok := b.dup[key]
	if !ok {
		return false
	}
	b.len--
	d.Delete(x)
	b.bst.replaceValue(key, d.Peek())
	if len(d) > 0 {
		return false
	}
	return b.bst.Delete(key)
}

func (b *bstDup) Predecessor(key int) interface{} {
	return b.bst.Predecessor(key)
}

func (b *bstDup) MinK() (int, interface{}) {
	return b.bst.MinK()
}
func (b *bstDup) MaxK() (int, interface{}) {
	return b.bst.MaxK()
}
func (b *bstDup) Min() interface{} {
	return b.bst.Min()
}
func (b *bstDup) Max() interface{} {
	return b.bst.Max()
}
func (b *bstDup) Len() int {
	return b.len
}
func (b *bstDup) Height() int {
	return b.bst.Height()
}
func (b *bstDup) IsValid() bool {
	return b.bst.IsValid()
}
func (b *bstDup) Slice() []interface{} {
	return b.Slice()
}

func (b *bstDup) String() string {
	return b.bst.String()
}
