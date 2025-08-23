package main

import (
	"flag"
	"os"
	"rosalind/utils"
	"strings"
)

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

	return res
}

func hamiltonianPath(edges [][]int) []int {
	// Finds if a Hamiltonian path exists within a DAG. Works by first
	// topologically sorting the DAG, then checking if there is an edge between
	// all adjacent nodes in the topologically-sorted version.
	order := topologicalSort(edges)
	res := make([]int, len(edges)+1)
	res[0] = 1

	adjMat := make([][]int, len(edges))
	for i := range adjMat {
		adjMat[i] = make([]int, len(edges))
	}
	for u, uEdges := range edges {
		for _, v := range uEdges {
			adjMat[u][v] = 1
		}
	}

	for i, v := range order[1:] {
		u := order[i]

		if adjMat[u][v] == 1 {
			res[i+1] = u + 1
			res[i+2] = v + 1
		} else {
			return []int{-1}
		}
	}

	return res
}

func main() {
	file := flag.String("file", "data.txt", "the file to be parsed as a DAGs.")
	flag.Parse()

	dat, err := os.ReadFile(*file)
	if err != nil {
		panic(err)
	}
	graphs := strings.Split(strings.TrimSpace(string(dat)), "\n\n")[1:]

	for _, graph := range graphs {
		edges := utils.EdgeListToDirAdjList(strings.Split(graph, "\n"))
		utils.RosalindPrintArr(hamiltonianPath(edges))
	}
}
