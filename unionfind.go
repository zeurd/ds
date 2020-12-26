package ds

//parent rank contain the pointer to parent and the rank
type parentRank struct {
	pointer int
	rank    int
}

// UnionFind is a a union-finds set
type UnionFind map[int]*parentRank

// NewUnionFind foo
func NewUnionFind() UnionFind {
	return make(UnionFind)
}

// Add adds the given element, each in their own new component
func (u UnionFind) Add(xs ...int) {
	for _, x := range xs {
		u[x] = &parentRank{x, 0}
	}
}

// Count returns the number of component
func (u UnionFind) Count() int {
	count := 0
	for x := range u {
		if u.isRoot(x) {
			count++
		}
	}
	return count
}

func (u UnionFind) isRoot(x int) bool {
	return u[x].pointer == x
}

// Find returns the group that x belongs to or -1
func (u UnionFind) Find(x int) int {
	pr, ok := u[x]
	if !ok {
		return -1
	}
	if u.isRoot(pr.pointer) {
		return pr.pointer
	}
	p := u.Find(pr.pointer)

	//path compression
	pr.pointer = p

	return p
}
func (u UnionFind) find(x int) (int, *int) {
	pr, ok := u[x]
	if !ok {
		return -1, nil
	}
	//fmt.Printf("path from %d: %d\n", x, pr.pointer)
	if u.isRoot(pr.pointer) {
		return x, &pr.rank
	}
	p, r := u.find(pr.pointer)
	//path compression
	pr.pointer = p
	return p, r
}

// Connected returns true if x and y belong to the same component
func (u UnionFind) Connected(x, y int) bool {
	return u.Find(x) == u.Find(y)
}

// Union unites 2 components
func (u UnionFind) Union(x, y int) {
	s1, r1 := u.find(x)
	s2, r2 := u.find(y)
	//- if rank[s1] > rank[s2]: set parent[s2] to s1
	if *r1 > *r2 {
		u[s2].pointer = s1
	} else {
		u[s1].pointer = s2
		if *r1 == *r2 {
			*r2++
		}
	}
}

// Rank returns the rank of an object
func (u UnionFind) Rank(x int) int {
	return u[x].rank
}
