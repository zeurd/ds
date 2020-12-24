package ds

import (
	"fmt"
	"math"
)

/*
TODO:
when order can take duplicates,
remove the duplicate count map d
as order could then be used to keep track of duplicates as well
*/

//heap is a min-heap
type heap struct {
	e  []element               // elements in the heap
	ko map[interface{}]*Order  // map[element.value] to Order containing the element.key: identical element.value can have different element.key
	p  map[element]int         // positions of elements in the heap; keys are element because one pair, key:value can only have one position
	d  map[element]int         // duplicate count for element that have both key and value indentical (all those identical values have same position in the heap)
	pk func(a interface{}) int // pk returns priority key for the given value a
}

type element struct {
	key   int
	value interface{}
}

func (h *heap) String() string {
	return fmt.Sprintf("%v", h.filterValues(h.e))
}

// NewHeap returns a heap without a specific eval function
func newHeap() *heap {
	return &heap{
		e:  make([]element, 0),
		ko: make(map[interface{}]*Order),
		p:  make(map[element]int),
		d:  make(map[element]int),
		pk: func(a interface{}) int { return a.(int) },
	}
}

// NewHeapWithEval foo
func newHeapWithEval(f func(a interface{}) int) *heap {
	return &heap{
		e:  make([]element, 0),
		ko: make(map[interface{}]*Order),
		p:  make(map[element]int),
		d:  make(map[element]int),
		pk: f,
	}
}

// Len gives the length of the heap (this does not count the perfect duplicates with same key values)
func (h *heap) Len() int {
	return len(h.e)
}

// Count counts element in the heap, duplicated included
func (h *heap) Count() int {
	c := 0
	for _, v := range h.d {
		c += v
	}
	return c
}

// Push adds an element to the heap and maintains its invariant
// It works only if the eval function has been implemented
func (h *heap) Push(e interface{}) {
	h.Insert(e, h.pk(e))
}

func (h *heap) swap(i, j int) {
	h.e[i], h.e[j] = h.e[j], h.e[i]
	h.p[h.e[i]], h.p[h.e[j]] = i, j
}

// parent return the parent of the element at the given position
func (h *heap) parent(i int) int {
	if i == 0 {
		return -1
	}
	if i%2 == 0 {
		return (i / 2) - 1
	}
	return i / 2
}

// left retuns the left child or - 1 if there's none
func (h *heap) left(i int) int {
	l := (i * 2) + 1
	if l >= h.Len() {
		return -1
	}
	return l
}

func (h *heap) right(i int) int {
	r := (i * 2) + 2
	if r >= h.Len() {
		return -1
	}
	return r
}

//smallerChild return the position of the smaller of the 2 children
func (h *heap) smallerChild(i int) int {
	l := h.left(i)
	r := h.right(i)
	if l == -1 {
		return -1 //no child
	}
	if r == -1 {
		return l //no right child
	}
	if h.e[l].key < h.e[r].key {
		return l
	}
	return r
}

func (h *heap) valid(i int) bool {
	return i < h.Len() && i >= 0
}

func (h *heap) IsValid() {
	if h.IsEmpty() {
		return
	}
	if h.Len() != len(h.d) || h.Len() != len(h.p) {
		s := fmt.Sprintf("heap length: %d, duplicates count: %d, postions: %d\n", h.Len(), len(h.d), len(h.p))
		panic(s)
	}
	//checks that min is top element
	min := h.ko[h.e[0].value].Min()
	for i, e := range h.e {
		if e.key < min {
			s := fmt.Sprintf("Position 0 key: %d. Postion %d key: %d\n", min, i, e.key)
			panic(s)
		}
	}
	//checks all parent-child relation
	for i := range h.e {
		h.checkParentRelation(i)
	}
}

func (h *heap) checkParentRelation(i int) {
	p := h.parent(i)
	if p != -1 && h.e[p].key > h.e[i].key {
		s := fmt.Sprintf("%v is parent of %v", h.e[p], h.e[i])
		panic(s)
	}
}

//Insert adds the element to the heap using the provided value to define priority
func (h *heap) Insert(x interface{}, v int) {
	// 1. stick x at end of last level
	e := element{v, x}
	h.e = append(h.e, e)
	if k, ok := h.ko[x]; ok{
		k.Add(v)
	} else {
		h.ko[x] = NewOrderFromInts(v)
	}
	//if an element with the same key:value exists, it's already at right position in the heap
	//just need to increase its duplicate count
	_, ok := h.d[e]
	if ok {
		h.d[e]++
		return
	}
	h.d[e] = 1
	h.p[e] = h.Len() - 1
	//2. bubble-up: if this violates the parent/child ruled: swap with parent, if this violates again, swap again
	h.bubbleUp(h.Len() - 1)
}

func (h *heap) bubbleUp(i int) {
	p := h.parent(i)
	if !h.valid(p) {
		return
	}
	if h.e[p].key > h.e[i].key {
		h.swap(p, i)
		h.bubbleUp(p)
	}
}

//Peek returns the top element without removing it
func (h *heap) Peek() interface{} {
	return h.e[0].value
}

//Pop returns the top element and removes it
func (h *heap) Pop() interface{} {
	last := h.Len() - 1
	if last == -1 {
		return nil
	}
	// 0. Min is at the root
	element := h.e[0]
	// 2. move last node to be new root
	h.swap(0, last)
	// 1. remove the root that has just been moved to last
	h.e = h.e[:last]
	if h.Len() == 1 {
		h.deleteInMaps(element)
		return element
	}
	// 3. bubble-down: swap new root => swap with smaller child, if tree still not ok, repeat the swap with smaller child (until there' no child)
	h.bubbleDown(0)
	h.deleteInMaps(element)
	return element.value
}

func (h *heap) delete(i int) {
	if h.IsEmpty() {
		return
	}
	end := h.Len() - 1
	if i == end {
		h.deleteLastElement()
		return
	}
	if i == end-1 && h.haveSameParent(i, end) {
		h.deleteBeforeLastElement()
		return
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
	return
}

func (h *heap) haveSameParent(i, j int) bool {
	return h.parent(i) == h.parent(j)
}

func (h *heap) deleteBeforeLastElement() {
	last := h.Len() - 1
	h.swap(last, last-1)
	h.deleteLastElement()
}

func (h *heap) deleteLastElement() {
	last := h.Len() - 1
	e := h.e[last]
	h.e = h.e[:last]
	h.deleteInMaps(e)
}

// TOFIX : logic to delete in keys
func (h *heap) deleteInMaps(e element) {
	if _, ok := h.ko[e.value]; !ok {
		//element was not present
		return
	}
	// decrease the duplicate count
	h.d[e]--
	if h.d[e] == 0 {
		delete(h.d, e)
		delete(h.p, e)
		h.ko[e.value].Delete(e.key)
		if h.ko[e.value].IsEmpty() {
			delete(h.ko, e.value)
		}
	}
}
func (h *heap) bubbleDown(i int) {
	c := h.smallerChild(i)
	if !h.valid(c) {
		return
	}
	if h.e[c].key < h.e[i].key {
		h.swap(c, i)
		h.bubbleDown(c)
	}
}

//IsEmpty returns true if the heap is empty
func (h *heap) IsEmpty() bool {
	return h.Len() == 0
}

//Contains returns true if heap contains given element
func (h *heap) Contains(x interface{}) bool {
	_, ok := h.ko[x]
	return ok
}

// Delete deletes the given element (if duplicates: the one with highest key, lowest priority)
func (h *heap) Delete(x interface{}) {
	if !h.Contains(x) {
		return
	}
	k := h.ko[x].Max()
	e := element{k, x}
	p := h.p[e]
	h.delete(p)
}

func (h *heap) DeleteKeyValue(k int, x interface{}) {
	e := element{k, x}
	p := h.p[e]
	h.delete(p)
}

// Update updates the priority value of the given element (if duplicates: the one with highest key, lowest priority)
func (h *heap) Update(x interface{}, key int) {
	keys, ok := h.ko[x]
	if !ok {
		return
	}
	k := keys.Max()
	h.UpdateKeyValue(x, k, key)
}

func (h *heap) UpdateKeyValue(x interface{}, oldKey, newKey int) {
	e := element{oldKey, x}
	pos := h.pos(e)
	if pos != -1 {
		h.delete(pos)
		h.Insert(x, newKey)
	}
}

// Pos returns the position of x in the heap or -1 if x is not present
func (h *heap) pos(e element) int {
	if v, ok := h.p[e]; ok {
		return v
	}
	return -1
}

//Value returns the value and if the element is present (returns the lowest value if duplicates)
func (h *heap) Value(x interface{}) (int, bool) {
	if v, ok := h.ko[x]; ok {
		return v.Min(), ok
	}
	return 0, false
}

// Levels returns the number of Levels in the tree
func (h *heap) Levels() int {
	return h.level(h.Len() - 1)
}

// GetLastElementPosInSubTree returns the postion in heap of the last element with root at position i
func (h *heap) GetLastElementPosInSubTree(i int) int {
	l := h.Len()
	if i == 0 {
		return l - 1
	}
	_, _, pos := h.getLastLevelInSubtree(i)
	return pos
}

// Subtree returns the subtree with element at i as root
func (h *heap) Subtree(i int) [][]interface{} {
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

func (h *heap) getLastLevelInSubtree(i int) ([]interface{}, int, int) {
	l := h.Levels()
	lastSublevel, s, e := h.subtreeLevel(i, l)
	if lastSublevel == nil || len(lastSublevel) == 0 {
		lastSublevel, s, e = h.subtreeLevel(i, l-1)
	}
	return lastSublevel, s, e
}

func (h *heap) level(i int) int {
	return int(math.Log2(float64(i + 1)))
}

// levelDetails returns the details about the level of the i position: level, start index, length of level
func (h *heap) levelDetailsAtPos(i int) (int, int, int) {
	level := h.level(i)
	start := int(math.Exp2(float64(level))) - 1
	finish := start + int(math.Exp2(float64(level)))
	return level, start, (finish - start)
}

// levelDetails returns the details about the l'th level: start index and length
func (h *heap) levelDetails(l int) (int, int) {
	start := int(math.Exp2(float64(l))) - 1
	finish := start + int(math.Exp2(float64(l)))
	return start, (finish - start)
}

// subtreeLevel returns the l'th subtree level of element at position i and its starting and ending index in heap
func (h *heap) subtreeLevel(i, l int) ([]interface{}, int, int) {
	iLevel, _, len := h.levelDetailsAtPos(i)
	diffLevel := l - iLevel
	if diffLevel == 0 {
		return h.filterValues(h.e[i : i+1]), i, i
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
	return h.filterValues(h.e[subtreeStart:subtreeEnd]), subtreeStart, subtreeEnd - 1
}

func (h *heap) filterValues(elements []element) []interface{} {
	v := make([]interface{}, len(elements))
	for i, e := range elements {
		v[i] = e.value
	}
	return v
}
