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
