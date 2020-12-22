package ds

type vertex int
type weight int

// Graph representation as map[vertex][vertex]weight
type Graph map[int]map[int]int

// Edge representation as {from, to, weight}
type Edge [3]int

//From returns the starting vertex of the Edge
func (e Edge) From() int {
	return e[0]
}

//To returns the ending vertex of the Edge
func (e Edge) To() int {
	return e[1]
}

// Weight returns the weight/length of the Edge
func (e Edge) Weight() int {
	return e[2]
}

//AddVertex adds a vertex to the graph
func (g Graph) AddVertex(v int) {
	g[v] = make(map[int]int)
}

// RemoveVertex removes the given vertex from the graph
func (g Graph) RemoveVertex(v int) {
	fromV := g[v]
	for _, w := range fromV {
		delete(g[w], v)
	}
	delete(g, v)
}

//AddNodesAndEdge adds both vertices of the given edge and the edge itself to the graph
func (g Graph) AddNodesAndEdge(e Edge) {
	g.AddVertex(e.From())
	g.AddVertex(e.To())
	g.PutEdge(e)
}

//AddNodesAndUndirectedEdge adds both vertices of the given edge and the edge itself to the graph
func (g Graph) AddNodesAndUndirectedEdge(e Edge) {
	g.AddVertex(e.From())
	g.AddVertex(e.To())
	g.PutUndirectedEdge(e)
}


// PutEdge adds or replace if it exists the given edge to the graph
// it panics if it adds an edge between unexisting node
func (g Graph) PutEdge(e Edge) {
	g[e.From()][e.To()] = e.Weight()
}

// PutUndirectedEdge adds or replace if it exists the given edge to the graph
// it panics if it adds an edge between unexisting node
func (g Graph) PutUndirectedEdge(e Edge) {
	g.PutEdge(e)
	g.PutEdge(Edge{e.To(), e.From(), e.Weight()})
}

// RemoveEdge removes the given edge if it exists
func (g Graph) RemoveEdge(e Edge) {
	delete(g[e.From()], e.To())
}

// RemoveUndirectedEdge removes the given undirected edge
func (g Graph) RemoveUndirectedEdge(e Edge) {
	delete(g[e.From()], e.To())
	delete(g[e.To()], e.From())
}

// ShortestPath implements dijkstra to return the shortest path from s to goal and its length
func (g Graph) ShortestPath(s, goal int) (int, []int) {
	X := NewSet()
	VX := NewHeap()
	A := make(map[int]int)
	B := make(map[int]int)

	for v := range g {
		VX.Insert(v, 1000000)
	}

	X.Add(s)
	VX.Update(s, 0)
	A[s] = 0
	B[s] = 1

	for {
		if VX.IsEmpty() {
			return -1, nil
		}
		v := VX.Pop().(int)

		if v == goal {
			return A[v], nil //g.path(B, s, goal)
		}

		for w, Lvw := range g[v] {
			X.Add(w)
			score := A[v] + Lvw
			if score < VX.Value(w) {
				A[w] = score
				B[w] = v
				VX.Update(w, score)
				g.checkWedges(w, X, VX, A, B)
			}
		}
	}
}

func (g Graph) checkWedges(w int, X Set, VX *Heap, A, B map[int]int) {
	for x, Lwx := range g[w] {
		if Lwx == 0 {
			panic(g[w])
		}
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

func (g Graph) path(B map[int]int, s, goal int) []int {
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
