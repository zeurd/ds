package ds

import (
	"fmt"
	"math"
)


//Heap is a min-heap
type Heap struct {
	e    []interface{}           // elements in the heap
	v    map[interface{}]int     // values to define priority (min)
	p    map[interface{}]int     // positions of elements in the heap
	eval func(a interface{}) int // eval gives the min value of the element
}

func (h *Heap) String() string {
	return fmt.Sprintf("%v", h.Subtree(0))
}

// NewHeap returns a heap without a specific eval function
func NewHeap() *Heap {
	return &Heap{
		e:    make([]interface{}, 0),
		v:    make(map[interface{}]int),
		p:    make(map[interface{}]int),
		eval: func(a interface{}) int { return a.(int) },
	}
}

// NewHeapWithEval foo
func NewHeapWithEval(f func(a interface{}) int) *Heap {
	return &Heap{
		e:    make([]interface{}, 0),
		v:    make(map[interface{}]int),
		p:    make(map[interface{}]int),
		eval: f,
	}
}

// Len gives the length of the heap
func (h *Heap) Len() int {
	return len(h.e)
}

// Push adds an element to the heap and maintains its invariant
// It works only if the eval function has been implemented
func (h *Heap) Push(e interface{}) {
	h.Insert(e, h.eval(e))
}

func (h *Heap) swap(i, j int) {
	h.e[i], h.e[j] = h.e[j], h.e[i]
	h.p[h.e[i]], h.p[h.e[j]] = i, j
}

// parent return the parent of the element at the given position
func (h *Heap) parent(i int) int {
	if i == 0 {
		return -1
	}
	if i%2 == 0 {
		return (i / 2) - 1
	}
	return i / 2
}

// left retuns the left child or - 1 if there's none
func (h *Heap) left(i int) int {
	l := (i * 2) + 1
	if l >= h.Len() {
		return -1
	}
	return l
}

func (h *Heap) right(i int) int {
	r := (i * 2) + 2
	if r >= h.Len() {
		return -1
	}
	return r
}

//smallerChild return the position of the smaller of the 2 children
func (h *Heap) smallerChild(i int) int {
	l := h.left(i)
	r := h.right(i)
	if l == -1 {
		return -1 //no child
	}
	if r == -1 {
		return l //no right child
	}
	if h.ValueAt(l) < h.ValueAt(r) {
		return l
	}
	return r
}

func (h *Heap) valid(i int) bool {
	return i < h.Len() && i >= 0
}

func (h *Heap) check() {
	if h.IsEmpty() {
		return
	}
	if h.Len() != len(h.v) || h.Len() != len(h.p) {
		s := fmt.Sprintf("Heap length: %d, Values: %d, Postions: %d\n%v\npos: %v\nval: %v", h.Len(), len(h.v), len(h.p), h, h.p, h.v)
		panic(s)
	}
	min := h.v[h.e[0]]
	for i, e := range h.e {
		if h.v[e] < min {
			s := fmt.Sprintf("At %d: %v with value: %d\nAt %d: %v with value: %d\n%v\n", 0, h.e[0], h.ValueAt(0), i, h.e[i], h.ValueAt(i), h)
			panic(s)
		}
	}
	for i := range h.e {
		h.checkParentRelation(i)
	}
}

func (h *Heap) checkParentRelation(i int) {
	p := h.parent(i)
	if p != -1 && h.ValueAt(p) > h.ValueAt(i) {
		s := fmt.Sprintf("%v with value %d is parent of %v with value %d", h.e[p], h.ValueAt(p), h.e[i], h.ValueAt(i))
		panic(s)
	}
}

//Insert adds the element to the heap using the provided value to define priority
func (h *Heap) Insert(x interface{}, v int) {
	// 1. strick j at end of last level
	h.e = append(h.e, x)
	h.v[x] = v
	h.p[x] = h.Len() - 1
	//2. bubble-up: if this violates the parent/child ruled: swap with parent, if this violates again, swap again
	h.bubbleUp(h.Len() - 1)
	h.check()
}

func (h *Heap) bubbleUp(i int) {
	p := h.parent(i)
	if !h.valid(p) {
		return
	}
	if h.v[h.e[p]] > h.v[h.e[i]] {
		h.swap(p, i)
		h.bubbleUp(p)
	}
}

//Peek returns the top element without removing it
func (h *Heap) Peek() interface{} {
	return h.e[0]
}

//Pop returns the top element and removes it
func (h *Heap) Pop() interface{} {
	return h.deletePop(0)
}

// deletePop is used when last leaf if subtree of i is the last leaf in global tree
func (h *Heap) deletePop(i int) interface{} {
	last := h.Len() - 1
	if last == -1 {
		return nil
	}
	// 0. Min is at the root
	element := h.e[i]
	// 2. move last node to be new root
	h.swap(i, last)
	// 1. remove the root that has just been moved to last
	h.e = h.e[:last]
	if h.Len() == 1 {
		h.deleteInMaps(element)
		return element
	}
	// 3. bubble-down: swap new root => swap with smaller child, if tree still not ok, repeat the swap with smaller child (until there' no child)
	h.bubbleDown(i)
	h.deleteInMaps(element)
	h.check()
	return element
}

func (h *Heap) delete(i int) interface{} {
	if h.IsEmpty() {
		return nil
	}
	end := h.Len() - 1
	if i == end {
		return h.deleteLastElement()
	}
	if i == end-1 && h.haveSameParent(i, end) {
		return h.deleteBeforeLastElement()
	}
	// in this case, we change subtree, so the last element that replaces the deleted one
	// might be smaller than its parent => need to try both bubble up and bubble down
	last := h.Len() - 1
	element := h.e[i]
	h.swap(i, last)
	h.e = h.e[:last]
	h.bubbleUp(i)
	h.bubbleDown(i)
	h.deleteInMaps(element)
	h.check()
	return element
}

func (h *Heap) haveSameParent(i, j int) bool {
	return h.parent(i) == h.parent(j)
}

func (h *Heap) deleteBeforeLastElement() interface{} {
	last := h.Len() - 1
	h.swap(last, last-1)
	return h.deleteLastElement()
}

func (h *Heap) deleteLastElement() interface{} {
	last := h.Len() - 1
	e := h.e[last]
	h.e = h.e[:last]
	h.deleteInMaps(e)
	h.check()
	return e
}

func (h *Heap) deleteInMaps(x interface{}) {
	delete(h.v, x)
	delete(h.p, x)
}
func (h *Heap) bubbleDown(i int) {
	c := h.smallerChild(i)
	if !h.valid(c) {
		return
	}
	if h.v[h.e[c]] < h.v[h.e[i]] {
		h.swap(c, i)
		h.bubbleDown(c)
	}
}

//IsEmpty returns true if the heap is empty
func (h *Heap) IsEmpty() bool {
	return h.Len() == 0
}

//Contains returns true if heap contains given element
func (h *Heap) Contains(x interface{}) bool {
	_, ok := h.v[x]
	return ok
}

//Delete deletes the given element if it's present in the heap
func (h *Heap) Delete(x interface{}) {
	if !h.Contains(x) {
		return
	}
	p := h.p[x]
	h.delete(p)
}

// Update updates the priority value of the given element
func (h *Heap) Update(x interface{}, value int) {
	// if h.Contains(x) && h.Value(x) < value {
	// 	panic(fmt.Sprintf("update: %v\n", x))
	// }
	pos := h.Pos(x)
	if pos != -1 {
		h.delete(pos)
		h.Insert(x, value)
	}
	h.check()
}

//Slice returns the heap as slice
func (h *Heap) Slice() []interface{} {
	return h.e
}

// Pos returns the position of x in the heap or -1 if x is not present
func (h *Heap) Pos(x interface{}) int {
	if v, ok := h.p[x]; ok {
		return v
	}
	return -1
}

//Value returns the value (min priority) or -1 if is not present
func (h *Heap) Value(x interface{}) int {
	if v, ok := h.v[x]; ok {
		return v
	}
	return -1
}

// Levels returns the number of Levels in the tree
func (h *Heap) Levels() int {
	return h.level(h.Len() - 1)
}

//ValueAt return the value of the
func (h *Heap) ValueAt(i int) int {
	x := h.e[i]
	return h.v[x]
}

// GetLastElementPosInSubTree returns the postion in heap of the last element with root at position i
func (h *Heap) GetLastElementPosInSubTree(i int) int {
	l := h.Len()
	if i == 0 {
		return l - 1
	}
	_, _, pos := h.getLastLevelInSubtree(i)
	return pos
}

// Subtree returns the subtree with element at i as root
func (h *Heap) Subtree(i int) [][]interface{} {
	start := h.level(i)
	end := h.level(h.Len() - 1)
	s := make([][]interface{}, 0)
	for l := start; l <= end; l++ {
		sub, _, _ := h.subtreeLevel(i, l)
		if len(sub) > 0 {
			s = append(s, sub)
		}
	}
	return s
}

func (h *Heap) getLastLevelInSubtree(i int) ([]interface{}, int, int) {
	l := h.Levels()
	lastSublevel, s, e := h.subtreeLevel(i, l)
	if lastSublevel == nil || len(lastSublevel) == 0 {
		lastSublevel, s, e = h.subtreeLevel(i, l-1)
	}
	return lastSublevel, s, e
}

func (h *Heap) level(i int) int {
	return int(math.Log2(float64(i + 1)))
}

// levelDetails returns the details about the level of the i position: level, start index, length of level
func (h *Heap) levelDetailsAtPos(i int) (int, int, int) {
	level := h.level(i)
	start := int(math.Exp2(float64(level))) - 1
	finish := start + int(math.Exp2(float64(level)))
	return level, start, (finish - start)
}

// levelDetails returns the details about the l'th level: start index and length
func (h *Heap) levelDetails(l int) (int, int) {
	start := int(math.Exp2(float64(l))) - 1
	finish := start + int(math.Exp2(float64(l)))
	return start, (finish - start)
}

// subtreeLevel returns the l'th subtree level of element at position i and its starting and ending index in heap
func (h *Heap) subtreeLevel(i, l int) ([]interface{}, int, int) {
	iLevel, _, len := h.levelDetailsAtPos(i)
	diffLevel := l - iLevel
	if diffLevel == 0 {
		return h.e[i : i+1], i, i
	}
	if diffLevel < 0 {
		return nil, -1, -1
	}
	iPosInLevel := (i + 1) % len
	subtreeLevelStart, _ := h.levelDetails(l)
	if subtreeLevelStart >= h.Len() {
		return nil, -1, -1
	}
	offset := int(math.Exp2(float64(diffLevel)))
	subtreeStart := subtreeLevelStart + (iPosInLevel * offset)
	if subtreeStart >= h.Len() {
		return nil, -1, -1
	}
	_, subtreeLen := h.levelDetails(diffLevel)
	subtreeEnd := subtreeStart + subtreeLen
	if subtreeEnd > h.Len() {
		subtreeEnd = h.Len()
	}
	return h.e[subtreeStart:subtreeEnd], subtreeStart, subtreeEnd - 1
}
