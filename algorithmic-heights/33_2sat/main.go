package main

import (
	"flag"
	"fmt"
	"os"
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

	return res
}


func connectedComponents(order []int, edges [][]int) []int {
	// Initialize all to discovered
	discovered := make([]bool, len(edges))
	components, c := make([]int, len(edges)), 0


	var dfs func(node int, edges [][]int, discovered []bool)
	dfs = func(node int, edges [][]int, discovered []bool) {
		if discovered[node] {
			return
		}
		discovered[node] = true
		components[node] = c

		for _, j := range edges[node] {
			if !discovered[j] {
				dfs(j, edges, discovered)
			}
		}
	}

	// For each node, DFS to identify component number
	for _, v := range order {
		// If this node already has an identified component, skip
		if discovered[v] {
			continue
		}

		dfs(v, edges, discovered)

		c++
	}

	return components
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


// Return components, component graph, and topologically sorted order in which
// to visit each strongly connected component in the resulting DAG
func stronglyConnectedComponents(graph [][]int) ([]int, [][]int, []int) {
	// Get the order of the reversed graph to find the post value order in which
	// to run the connected components algorithm on.
	graphR := reverseEdgeList(graph)
	order := topologicalSort(graphR)

	// Create strongly connected comps graph
	comps := connectedComponents(order, graph)
	compG := componentsToEdgeList(comps, graph)

	// Topologically sort components graph to identify sinks
	compOrder := topologicalSort(reverseEdgeList(compG))
	return comps, compG, compOrder
}


// Converts a list of strings in 2SAT edge list format to a directed adjacency
// list. 2SAT format is `x y` where `x`, `y` can be +ve or -ve. Therefore, we
// need to make 2*n nodes.
//
// To correct for negative node labels, I add n to negative labels and (n-1) to
// positive labels. This means for a node in position i, the corresponding -ve
// label is in position (2n-i-1)
func twoSATParse(arr []string) (int, [][]int) {
	n, _ := strconv.Atoi(strings.Split(arr[0], " ")[0]) // nodes

	edges := make([][]int, n*2)
	for _, edge := range arr[1:] {
		nodes := strings.Split(edge, " ")
		n1, _ := strconv.Atoi(nodes[0])
		n2, _ := strconv.Atoi(nodes[1])
		if n1 < 0 {
			n1 += n
		} else {
			n1 += n - 1
		}
		if n2 < 0 {
			n2 += n
		} else {
			n2 += n - 1
		}

		// for n1 V n1 :: (1) -n1 -> n2; (2) -n2 -> n1
		edges[2*n - n1 - 1] = append(edges[2*n - n1 - 1], n2)
		edges[2*n - n2 - 1] = append(edges[2*n - n2 - 1], n1)
	}

	return n, edges
}

// Solves a 2SAT (Rosalind-formatted) graph in edge list format by (1) finding
// the strongly connected components (SSCs) of the graph; and (2) identifying
// whether any literal's negation is contained within the same component (which
// indicates unsatisfiability); or (3) repeatedly choosing sinks of the
// component graph and setting contained literals to true.
//
// Returns (sol, true) if a solution is found; else (0, false)
func twoSAT(graph_s []string) ([]int, bool) {
	n, graph := twoSATParse(graph_s)
	// fmt.Println("Graph:", graph)

	// Get the SSCs + graph & order in which to identify sinks
	comps, compG, compO := stronglyConnectedComponents(graph)

	// Exit if any SSC contains a literal and its negation
	compNodes := make([][]int, len(compG)) // O(1) lookup of nodes in a component
	for i, c := range comps {
		if comps[2*n - i - 1] == c {
			return []int{}, false
		}
		compNodes[c] = append(compNodes[c], i)
	}

	// Otherwise, repeatedly visit sinks & assign true to literals
	sol := make([]int, 2*n)
	for _, c := range compO {
		for _, i := range compNodes[c] {
			// skip previously set literal values
			if sol[i] == 0 {
				// non-negated literals
				if i >= n {
					sol[i] = i - n + 1
					sol[2*n - i - 1] = -1
				} else {
					sol[i] = -1
					sol[2*n - i - 1] = -(2*n - i - 1 - n + 1)
				}
			}
		}
	}

	return sol[n:], true
}


func main() {
	// Read file
	file := flag.String("file", "data.txt", "the file to parse as a graph.")
	flag.Parse()

	dat, err := os.ReadFile(*file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n\n")

	// Find solutions for each
	for _, graph := range lines[1:] {
		sol, sat := twoSAT(strings.Split(graph, "\n"))
		if sat {
			fmt.Println("1", strings.Trim(fmt.Sprint(sol), "[]"))
		} else {
			fmt.Println("0")
		}
	}
}
