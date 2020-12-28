package ds

//Clusters is a UnionFind for interface
type Clusters struct {
	uf      UnionFind
	indices map[interface{}]int
	count   int
	i       int
}

// NewClusters returns a new Clusters
func NewClusters() *Clusters {
	return &Clusters{
		NewUnionFind(),
		make(map[interface{}]int),
		0,
		0,
	}
}

// Count returns the number of clusters
func (c *Clusters) Count() int {
	return c.count
}

// Add adds x to its own cluster
func (c *Clusters) Add(x interface{}) {
	c.count++
	c.i++
	c.indices[x] = c.i
	c.uf.Add(c.i)
}

// Find returns the cluster that x belongs to
func (c *Clusters) Find(x interface{}) int {
	ix := c.indices[x]
	return c.uf.Find(ix)
}

// Connected returns trye if x and y are in the same cluster
func (c *Clusters) Connected(x, y int) bool {
	return c.Find(x) == c.Find(y)
}

// Union brings merges the 2 cluster that contain x and y
func (c *Clusters) Union(x, y interface{}) {
	ix := c.indices[x]
	iy := c.indices[y]
	b := c.uf.Union(ix, iy)
	if b {
		c.count--
	}
}
