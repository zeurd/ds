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
			if !nProvided{
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
