package ds

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadGraph reads a txt file containing Graph data in format:
// [starting vertex] [ending vertex, edge length]...
func ReadGraph(file string) Graph {
	A := NewGraph()

	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		tk := s.Text()
		row := strings.Split(tk, "\t")
		startingNode := 0
		for j, edgestr := range row {
			if edgestr == "" {
				continue
			}
			//starting node
			if j == 0 {
				n, err := strconv.Atoi(edgestr)
				if err != nil {
					panic(edgestr)
				}
				startingNode = n
				continue
			}
			edgeDetails := strings.Split(edgestr, ",")
			if len(edgeDetails) != 2 {
				panic(edgestr)
			}
			toNode, err := strconv.Atoi(edgeDetails[0])
			weight, err := strconv.Atoi(edgeDetails[1])
			if err != nil {
				panic(edgestr)
			}
			A.AddVertex(startingNode)
			A.AddVertex(toNode)
			A.PutEdge(startingNode, toNode, weight)

		}
	}
	err = s.Err()
	if err != nil {
		fmt.Println(err)
	}
	return A
}

// ReadVE reads a txt file containing graph data in the form:
// first line: m n
// following lines: v w cost
func ReadVE(file string, undirected bool) (Graph, int, int, int) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return nil, 0, 0, 0
	}
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	g := NewGraph()
	s := bufio.NewScanner(f)
	i := 0
	var cost, m, n int
	nProvided := false
	for s.Scan() {
		if i == 0 {
			i++
			mn := strings.Split(s.Text(), " ")
			m, _ = strconv.Atoi(mn[0])
			if len(mn) > 1 {
				n, _ = strconv.Atoi(mn[1])
				nProvided = true
			}
			continue
		}
		job := strings.Split(s.Text(), " ")
		v, _ := strconv.Atoi(job[0])
		w, _ := strconv.Atoi(job[1])
		c, _ := strconv.Atoi(job[2])
		cost += c
		g.AddVertex(v)
		g.AddVertex(w)
		g.PutEdge(v, w, c)
		if !nProvided {
			n++
		}
		if undirected {
			g.PutEdge(w, v, c)
			if !nProvided {
				n++
			}
		}
	}
	err = s.Err()
	if err != nil {
		fmt.Println(err)
	}
	return g, m, n, cost
}

// ReadClustering reads a file in the following format:
// [# of nodes] [# of bits for each node's label]
// [first bit of node 1] ... [last bit of node 1]
// Use the Hamming distance--- the number of differing bits --- between the two nodes' labels.
// Eg the Hamming distance between the 24-bit label of node #2 above and the label "0 1 0 0 0 1 0 0 0 1 0 1 1 1 1 1 1 0 1 0 0 1 0 1" is 3
// (since they differ in the 3rd, 7th, and 21st bits)
// it returns the Graph, m, #bits in each label (= max distance)
func ReadClustering(file string) (Graph, int, int) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return nil, 0, 0
	}
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	g := NewGraph()
	s := bufio.NewScanner(f)
	i := 0
	vert := make(map[string]int)
	var m, bits int
	for s.Scan() {
		if i == 0 {
			i++
			mn := strings.Split(s.Text(), " ")
			m, _ = strconv.Atoi(mn[0])
			bits, _ = strconv.Atoi(mn[1])
			continue
		}
		bitsStr := strings.ReplaceAll(s.Text(), " ", "")
		v, _ := strconv.ParseInt(bitsStr, 2, 64)
		vert[bitsStr] = int(v)
		g.AddVertex(int(v))
	}
	err = s.Err()
	if err != nil {
		fmt.Println(err)
	}
	for v := range vert {
		at1 := atDistance(v, bits)
		for _, vat1 := range at1 {
			v1, ok := vert[vat1]
			if ok {
				g.PutEdge(vert[v], v1, 1)
			}
			at2 := atDistance(vat1, bits)
			for _, vat2 := range at2 {
				v2, ok := vert[vat2]
				if ok {
					g.PutEdge(vert[v], v2, 2)
				}
			}
		}
	}
	return g, m, bits
}

func hammingDistance(s1, s2 string) int {
	dist := 0
	for i, r := range s1 {
		if rune(s2[i]) != r {
			dist++
		}
	}
	return dist
}

func atDistance(s string, bits int) []string {
	b := []byte(s)
	ad := make([]string, bits)
	for i, c := range b {
		mod := []byte(s)
		if c == '1' {
			mod[i] = '0'
		} else {
			mod[i] = '1'
		}
		ad[i] = string(mod)
	}
	//panic(ad)
	return ad
}
