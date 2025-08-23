package main

import (
	"flag"
	"fmt"
	"os"
	"rosalind/utils"
	"slices"
	"strconv"
	"strings"
)

// Creates an adjacency list that is the reverse of the supplied edge list.
func reverseEdgeList(arr []string) [][]int {
	n, _ := strconv.Atoi(strings.Split(arr[0], " ")[0])

	edges := make([][]int, n)
	for _, edge := range arr[1:] {
		nodes := strings.Split(edge, " ")
		n1, _ := strconv.Atoi(nodes[1])
		n2, _ := strconv.Atoi(nodes[0])

		n1--
		n2--
		edges[n1] = append(edges[n1], n2)
	}

	return edges
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

	return res
}

func dfs(node int, edges [][]int, discovered []bool, component int, components []int) {
	if discovered[node] {
		return
	}
	discovered[node] = true
	components[node] = component

	for _, j := range edges[node] {
		if !discovered[j] {
			dfs(j, edges, discovered, component, components)
		}
	}
}

func connectedComponents(order []int, edges [][]int) []int {
	// Initialize all to discovered
	discovered := make([]bool, len(edges))
	components := make([]int, len(edges))
	component := 0

	// For each node, DFS to identify component number
	for _, v := range order {
		// If this node already has an identified component, skip
		if discovered[v] {
			continue
		}

		dfs(v, edges, discovered, component, components)

		component++
	}

	return components
}

func stronglyConnectedComponents(edgeList []string) ([][]int, []int) {
	graph := utils.EdgeListToDirAdjList(edgeList)
	graphR := reverseEdgeList(edgeList)

	// Get the order of the reversed graph to find the post value
	// order in which to run the connected components algorithm
	// on.
	order := topologicalSort(graphR)
	return graph, connectedComponents(order, graph)
}

func componentsToEdgeList(comps []int, edges [][]int) [][]int {
	// Finding max component
	maxComp := 0
	for _, c := range comps {
		if c > maxComp {
			maxComp = c
		}
	}
	maxComp++

	// Creating adj matrix
	adjMat := make([][]int, maxComp)
	for i := range adjMat {
		adjMat[i] = make([]int, maxComp)
	}

	for u := range edges {
		for _, v := range edges[u] {
			compU, compV := comps[u], comps[v]

			if compU == compV {
				continue
			}

			adjMat[compU][compV] = 1
		}
	}

	// Converting to edge list representation
	edgeList := make([][]int, maxComp)
	for u := range adjMat {
		for v := range adjMat[u] {
			if adjMat[u][v] == 1 {
				edgeList[u] = append(edgeList[u], v)
			}
		}
	}

	return edgeList
}

func semiConnected(graph []string) bool {
	// Solves the semi-connected problem in O(V + E) time (theoretically... my
	// use of graph representation is all over the place). First, computes the
	// strongly connected components of the graph, then topologically sorts the
	// resulting DAG (meta-graph). If there is an edge from each i-th to i+1-th
	// node in the sorted DAG, the graph is semi-connected.

	// Computing SCCs
	edges, comps := stronglyConnectedComponents(graph)
	metaEdges := componentsToEdgeList(comps, edges)

	// With the meta graph, topologically sort them
	order := topologicalSort(metaEdges)

	for i := range order[:len(order)-1] {
		if !slices.Contains(metaEdges[order[i]], order[i+1]) {
			return false
		}
	}

	return true
}

func main() {
	file := flag.String("file", "data.txt", "the file to parse as a graph.")
	flag.Parse()

	dat, err := os.ReadFile(*file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n\n")
	k, _ := strconv.Atoi(lines[0])

	res := make([]int, k)
	for i, graph := range lines[1:] {
		if semiConnected(strings.Split(graph, "\n")) {
			res[i] = 1
		} else {
			res[i] = -1
		}
	}

	fmt.Println(res)
}
