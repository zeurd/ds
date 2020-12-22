package ds

// Graph defines the graph interface
type Graph interface {
	Len() int
	AddVertex(v int)
	RemoveVertex(v int)
	PutEdge(from, to, weight int)
	//PutUndirectedEdge(from, to int)
	RemoveEdge(from, to int)
	//RemoveUndirectedEdge(from, to int)
	ShortestPath(from, to int) (int, []int)
	EdgesCost() (int, int)
	MST() Graph
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
