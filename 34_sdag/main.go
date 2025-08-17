package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"rosalind/utils"
	"strconv"
	"strings"
)

// Creates an adjacency list that is the reverse of the supplied edge list.
func reverseEdgeList(arr [][]int) [][]int {
	edgesR := make([][]int, len(arr))

	for n1, edges := range arr {
		for _, n2 := range edges {
			edgesR[n2] = append(edgesR[n2], n1)
		}
	}

	return edgesR
}


func topologicalSort(edges [][]int) []int {
	n := len(edges)

	discovered := make([]bool, n)
	res, i := make([]int, n), n-1

	var dfs func(u int)
	dfs = func(u int) {
		discovered[u] = true

		for _, v := range edges[u] {
			if !discovered[v] {
				dfs(v)
			}
		}
		res[i] = u
		i--
	}

	for u := range edges {
		if !discovered[u] {
			dfs(u)
		}
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}


/* Algorithm:
- Initialise all distances to inf (& source as 0)
- Topologically sort the graph
- Visit each node u in topological order, and for each adjacent vertex v:
	- if dist[v] > dist[u] + weight(u, v): dist[v] = dist[u] + weight(u, v)
*/
func main() {
	// Read file
	file := flag.String("file", "data.txt", "the file to parse as a graph.")
	flag.Parse()

	dat, err := os.ReadFile(*file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
	n, _ := strconv.Atoi(strings.Split(lines[0], " ")[0])

	g := utils.EdgeListToDirAdjList(lines)
	ws := utils.WeightsFromEdgeList(lines)

	// topologically sort DAG
	gR := reverseEdgeList(g)
	order := topologicalSort(gR)

	// distances
	ds := make([]int, n)
	for i := range ds {
		ds[i] = math.MaxInt
	}
	ds[0] = 0

	// traversing
	for _, u := range order {
		if ds[u] != math.MaxInt {
			for _, v := range g[u] {
				if ds[v] > ds[u] + ws[u][v] {
					ds[v] = ds[u] + ws[u][v]
				}
			}
		}
	}

  for _, d := range ds {
    if d == math.MaxInt {
      fmt.Print("x ")
    } else {
      fmt.Printf("%d ", d)
    }
  }
  fmt.Printf("\n")
}
