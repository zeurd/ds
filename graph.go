package ds

// Graph defines the graph interface
type Graph interface {
	Len() int
	Edges() Set
	Nodes() Set
	AddVertex(v int)
	RemoveVertex(v int)
	PutEdge(from, to, weight int)
	//PutUndirectedEdge(from, to int)
	RemoveEdge(from, to int)
	//RemoveUndirectedEdge(from, to int)
	ShortestPath(from, to int) (int, []int)
	EdgesCost() (int, int)
	MST() (Graph, int)
	Clusters(k int) int
	ClustersDist(d int) int
	//BFSP(s,m,n int) ([]map[int]int, int)
	AllPairsSP() interface{}
}

// NewGraph returns a default graph: directed and with int vertices
func NewGraph() Graph {
	return make(graph)
}

// NewUndirectedGraph

// NewTypedGraph(type, directed)

// Options struct?
// undirected bool
// type Type
// weighted bool
// negative wieghts
// preprocessed for A*
// parallel ?

//Edge is [node, weight]
type Edge [3]int

//From returns the origin vertex
func (e Edge) From() int {
	return e[0]
}

//To returns the destination vertex
func (e Edge) To() int {
	return e[1]
}

//Weight returns the length of the edge
func (e Edge) Weight() int {
	return e[2]
}
