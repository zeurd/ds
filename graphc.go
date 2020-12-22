package ds

// graph representation as map[int][int]int, for map[from]map[to]weight
type graph map[int]map[int]int

// Len returns the number of vertices in the graph
func (g graph) Len() int {
	return len(g)
}

//AddVertex adds a vertex to the graph
func (g graph) AddVertex(v int) {
	if _, ok := g[v]; !ok {
		g[v] = make(map[int]int)
	}
}

// RemoveVertex removes the given vertex from the graph
func (g graph) RemoveVertex(v int) {
	fromV := g[v]
	for _, w := range fromV {
		delete(g[w], v)
	}
	delete(g, v)
}

// PutEdge adds or replace if it exists the given edge to the graph
// it panics if it adds an edge between unexisting node
func (g graph) PutEdge(from, to, weight int) {
	if cost, ok := g[from][to]; ok && cost < weight {
		return
	}
	g[from][to] = weight
}

// PutUndirectedEdge adds or replace if it exists the given edge to the graph
// it panics if it adds an edge between unexisting node
func (g graph) PutUndirectedEdge(from, to, weight int) {
	g[from][to] = weight
	g[to][from] = weight
}

// RemoveEdge removes the given edge if it exists
func (g graph) RemoveEdge(from, to int) {
	delete(g[from], to)
}

// RemoveUndirectedEdge removes the given undirected edge
func (g graph) RemoveUndirectedEdge(from, to int) {
	delete(g[from], to)
	delete(g[to], from)
}

// EdgesCost returns the numbers of edges and the total sum of their weights
func (g graph) EdgesCost() (int, int) {
	cost := 0
	e := 0
	for _, edges := range g {
		for _, c := range edges {
			cost += c
			e++
		}
	}

	return e, cost
}

// ShortestPath implements dijkstra to return the shortest path from s to goal and its length
func (g graph) ShortestPath(from, to int) (int, []int) {
	X := NewSet()
	VX := NewHeap()
	A := make(map[int]int)
	B := make(map[int]int)

	for v := range g {
		VX.Insert(v, 1<<32-1)
	}

	X.Add(from)
	VX.Update(from, 0)
	A[from] = 0
	B[to] = 1

	for {
		if VX.IsEmpty() {
			return -1, nil
		}
		v := VX.Pop().(int)

		if v == to {
			return A[v], nil //g.path(B, s, goal)
		}

		for w, Lvw := range g[v] {
			X.Add(w)
			score := A[v] + Lvw
			val, ok := VX.Value(w)
			if ok && score < val {
				A[w] = score
				B[w] = v
				VX.Update(w, score)
				g.checkWedges(w, X, VX, A, B)
			}
		}
	}
}

func (g graph) checkWedges(w int, X Set, VX *Heap, A, B map[int]int) {
	for x, Lwx := range g[w] {
		if !X.Contains(x) {
			newScore := A[w] + Lwx
			oldScore, ok := A[x]
			if ok && oldScore <= newScore {
				return
			}
			VX.Update(x, newScore)
			A[x] = newScore
			B[x] = w
		}
	}
}

func (g graph) path(B map[int]int, s, goal int) []int {
	p := []int{goal}
	w := goal
	for {
		p = append(p, B[w])
		if B[w] == s {
			return p
		}
		w = B[w]
	}
}

func (g graph) MST() (Graph, int) {
	X := NewSet()
	VX := NewHeap()
	A := make(map[int]struct{ v, c int })
	T := NewGraph()
	total := 0
	for v := range g {
		VX.Insert(v, 1<<32-1)
	}

	s := 37
	X.Add(s)
	VX.Update(s, 0)
	A[s] = struct{ v, c int }{s, 0}

	for {
		if T.Len() == g.Len() {
			return T, total
		}
		v := VX.Pop().(int)
		T.AddVertex(v)
		for w, cost := range g[v] {
			X.Add(w)
			T.AddVertex(w)
			val, ok := VX.Value(w)
			if ok && cost < val {
				A[w] = struct{ v, c int }{v, cost}
				total += cost
				T.PutEdge(v, w, cost)
				T.PutEdge(w, v, cost)
				VX.Update(w, cost)
				g.checkCosts(w, X, VX, T, A, &total)
			}
		}
	}
	return T, 0
}

func (g graph) checkCosts(w int, X Set, VX *Heap, T Graph, A map[int]struct{ v, c int }, total *int) {
	for x, newCost := range g[w] {
		if !X.Contains(x) {
			vc, ok := A[x]
			oldCost := vc.c
			if ok && oldCost <= newCost {
				return
			}
			VX.Update(x, newCost)
			*total += newCost
			*total -= oldCost
			T.RemoveEdge(vc.c, w)
			T.RemoveEdge(w, vc.c)
			T.AddVertex(x)
			T.PutEdge(w, x, newCost)
			T.PutEdge(x, w, newCost)
			A[x] = struct{ v, c int }{w, newCost}
		}
	}
}
