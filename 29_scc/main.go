package main

import (
	"flag"
	"fmt"
	"rosalind/utils"
	"strconv"
	"strings"
)

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

func dfs(node int, edges [][]int, discovered []bool) {
	if discovered[node] {
		return
	}
	discovered[node] = true
	for _, j := range edges[node] {
		if !discovered[j] {
			dfs(j, edges, discovered)
		}
	}
}

func numConnectedComponents(order []int, edges [][]int) int {
	// Initialize all to discovered
	discovered := make([]bool, len(edges))
	components := 0

	// For each node, DFS to identify component number
	for _, v := range order {
		// If this node already has an identified component, skip
		if discovered[v] {
			continue
		}

		dfs(v, edges, discovered)

		components++
	}
	return components
}

func stronglyConnectedComponents(edgeList []string) int {
	graph := utils.EdgeListToDirAdjList(edgeList)
	graphR := reverseEdgeList(edgeList)

	// Get the order of the reversed graph to find the post value
	// order in which to run the connected components algorithm
	// on.
	order := topologicalSort(graphR)
	return numConnectedComponents(order, graph)
}

func main() {
	file := flag.String("file", "data.txt", "the file to parse as a graph.")
	flag.Parse()

	lines, _ := utils.ReadLines(*file)

	fmt.Println(stronglyConnectedComponents(lines))
}
