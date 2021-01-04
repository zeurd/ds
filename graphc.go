package ds

import (
	"fmt"
	"time"
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

//Edges returns a set containing all the edges of the graph
func (g graph) Edges() Set {
	set := NewSet()
	for v, edges := range g {
		for w, c := range edges {
			set.Add(Edge{v, w, c})
		}
	}
	return set
}

// Nodes returns a set containing all the nodes in the graph
func (g graph) Nodes() Set {
	set := NewSet()
	for n := range g {
		set.Add(n)
	}
	return set
}

// ShortestPath implements dijkstra to return the shortest path from s to goal and its length
func (g graph) ShortestPath(from, to int) (int, []int) {
	X := NewSet()
	VX := NewHeap(false)
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
			return A[v], g.path(B, from, to)
		}

		for w, Lvw := range g[v] {
			X.Add(w)
			score := A[v] + Lvw
			val, ok := VX.Key(w)
			if ok && score < val {
				A[w] = score
				B[w] = v
				VX.Update(w, score)
				g.checkWedges(w, X, VX, A, B)
			}
		}
	}
}

func (g graph) checkWedges(w int, X Set, VX Heap, A, B map[int]int) {
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

// MST returns the minimum spanning tree and its cost using Prim's alg
func (g graph) MST() (Graph, int) {
	X := NewSet()
	VX := NewHeap(false)
	A := make(map[int]Edge)
	T := NewGraph()
	total := 0

	s := 6
	X.Add(s)

	//preprocess
	for v, e := range g {
		if v == s {
			continue
		}
		// there is an edge (s,v)
		if cost, ok := e[s]; ok {
			VX.Insert(v, cost)
			A[v] = Edge{s, v, cost}
		} else {
			// there is no crossing edge
			VX.Insert(v, 1<<32-1)
		}
	}

	//main
	for !VX.IsEmpty() {
		w := VX.Pop().(int)
		X.Add(w)
		winner := A[w]
		T.AddVertex(winner.From())
		T.AddVertex(winner.To())
		T.PutEdge(winner.From(), winner.To(), winner.Weight())
		total += winner.Weight()

		for y, cost := range g[w] {
			if X.Contains(y) {
				continue
			}
			oldKey, ok := VX.Key(y)
			if !ok || ok && oldKey > cost {
				VX.Update(y, cost)
				A[y] = Edge{w, y, cost}
			}
		}
	}
	return T, total
}

// Clusters implement the Kruskal algo to compute Max-Spacing k-clusters and returns the spacing
func (g graph) Clusters(k int) int {
	return g.clustersB(k)
	closest := NewOrderedList(func(e interface{}) int { return e.(Edge).Weight() })
	clusters := NewClusters()
	// get the closest node to each others, and add each in its own clusters

	now := time.Now()
	for v, edges := range g {
		for w, c := range edges {
			closest.Add(Edge{v, w, c})
			// clusters.Add(v)
			// clusters.Add(w)
		}
	}
	now2 := time.Now()
	fmt.Printf("add  in order took: %v (%d elements)\n", now2.Sub(now), closest.Len())
	for v, edges := range g {
		for w := range edges {
			//closest.Add(Edge{v, w, c})
			clusters.Add(v)
			clusters.Add(w)
		}
	}
	now3 := time.Now()
	fmt.Printf("add in clusters: %v\n", now3.Sub(now2))
	// merge the clusters of the 2 closest nodes until k clusters
	var last int
	closestS := closest.Slice()
	for i, e := range closestS {
		edge := e.(Edge)
		clusters.Union(edge.From(), edge.To())
		if clusters.Count() == k {
			last = i
			break
		}
	}
	fmt.Printf("after first loop: %d elements in closest\n", closest.Len()-last)
	// find the next closest pair that is separated
	for i := last + 1; i < len(closestS); i++ {
		next := closest.Get(i).(Edge)
		if !clusters.Connected(next.From(), next.To()) {
			fmt.Printf("after second loop: %d elements in closest\n", closest.Len()-i)
			return next.Weight()
		}
	}
	return 0
}

func (g graph) clustersB(k int) int {
	closeBst := NewBinarySearchTree(true, func(x interface{}) int { return x.(Edge).Weight() })
	clusters := NewClusters()

	now := time.Now()
	for v, edges := range g {
		for w, c := range edges {
			closeBst.Push(Edge{v, w, c})
		}
	}
	now2 := time.Now()
	fmt.Printf("add  in order took: %v (%d elements)\n", now2.Sub(now), closeBst.Len())
	for v, edges := range g {
		for w := range edges {
			//closest.Add(Edge{v, w, c})
			clusters.Add(v)
			clusters.Add(w)
		}
	}
	now3 := time.Now()
	fmt.Printf("add in clusters: %v\n", now3.Sub(now2))
	// merge the clusters of the 2 closest nodes until k clusters

	for closeBst.Len() > 0 {
		k, e := closeBst.MinK()
		closeBst.Delete(k)
		edge := e.(Edge)
		clusters.Union(edge.From(), edge.To())
		if clusters.Count() == k {
			break
		}
	}
	fmt.Printf("after first loop: %d elements in closest\n", closeBst.Len())

	// find the next closest pair that is separated
	for closeBst.Len() > 0 {
		k, e := closeBst.MinK()
		closeBst.Delete(k)
		edge := e.(Edge)
		if !clusters.Connected(edge.From(), edge.To()) {
			fmt.Printf("after second loop: %d elements in closest\n", closeBst.Len())
			return edge.Weight()
		}
	}
	return 0
}

func (g graph) ClustersDist(d int) int {
	closest := NewOrderedList(func(e interface{}) int { return e.(Edge).Weight() })
	clusters := NewClusters()
	// get the closest node to each others, and add each in its own clusters
	for v, edges := range g {
		for w, c := range edges {
			closest.Add(Edge{v, w, c})
			clusters.Add(v)
			clusters.Add(w)
		}
	}
	fmt.Printf("closest size: %d\n", closest.Len())
	// merge the clusters of the 2 closest nodes until k clusters
	closestS := closest.Slice()
	for _, e := range closestS {

		edge := e.(Edge)
		if clusters.Connected(edge.From(), edge.To()) {
			continue
		}
		if edge.Weight() >= d {
			return clusters.Count()
		}
		clusters.Union(edge.From(), edge.To())
	}
	return clusters.Count()
}
