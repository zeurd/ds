package ds

import (
)

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
		VX.Insert(v, 1000000)
	}

	X.Add(from)
	VX.Update(from, 0)
	A[to] = 0
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
			if score < VX.Value(w) {
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
