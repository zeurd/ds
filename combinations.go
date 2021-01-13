package ds

import "fmt"

// Combinations foo
type Combinations struct {
	n            int
	combinations [][]int
	bitmasks     [][]uint8
}

// Size returns the number of all possible combinations
func (c *Combinations) Size() int {
	return len(c.combinations)
}

// NewCombinations foo
func NewCombinations(n, r int) *Combinations {
	c := &Combinations{
		n,
		make([][]int, 0, n),
		make([][]uint8, 0, n),
	}
	data := make([]int, r)
	c.combine(data, 0, n-1, 0)
	return c
}

func (c *Combinations) add(data []int) {
	c.combinations = append(c.combinations, data)
}

func (c *Combinations) combine(data []int, start, end, i int) {
	if i == len(data) {
		d := make([]int, len(data))
		bitmask := make([]uint8, c.n)
		copy(d, data)
		for _, x := range d {
			bitmask[x] = 1
		}
		c.bitmasks = append(c.bitmasks, bitmask)
		c.add(d)
		return
	}
	if start <= end {
		data[i] = start
		c.combine(data, start+1, end, i+1)
		c.combine(data, start+1, end, i)
	}
}

func (c *Combinations) String() string {
	return fmt.Sprintf("%v", c.combinations)
}

// All returns all the combinatins
func (c *Combinations) All() [][]int {
	return c.combinations
}

// Bitmasks foo
func (c *Combinations) Bitmasks() [][]uint8 {
	return c.bitmasks
}
