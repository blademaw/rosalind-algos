package main

import (
	"flag"
	"os"
	"rosalind/utils"
	"strconv"
	"strings"
)

func bellmanFordNegCycle(s int, edges [][]int, weights [][]int, dist []int, found []bool) bool {
	// Returns whether a graph has a negative cycle. We can find this by using
	// Bellman-Ford from an arbitrary source node and relaxing edges V-1 times,
	// and checking if when relaxing for the V-th time any distance estimates
	// improve. This means we have a negative cycle.

	for i := range found {
		found[i] = false
		dist[i] = 0
	}
	found[s] = true

	for i := 0; i < len(edges)-1; i++ {
		for u, uEdges := range edges {
			if !found[u] {
				continue
			}

			for _, v := range uEdges {
				if !found[v] || dist[v] > dist[u]+weights[u][v] {
					dist[v] = dist[u] + weights[u][v]
					found[v] = true
				}
			}
		}
	}

	// Relaxing for the V-th time
	for u, uEdges := range edges {
		if !found[u] {
			continue
		}

		for _, v := range uEdges {
			if !found[v] || dist[v] > dist[u]+weights[u][v] {
				return true
			}
		}
	}

	return false
}

func negCycle(edges [][]int, weights [][]int) bool {
	visited := make([]bool, len(edges))
	dist := make([]int, len(edges))
	found := make([]bool, len(edges))

	for s := range edges {
		if !visited[s] {
			// Try to find a cycle
			if bellmanFordNegCycle(s, edges, weights, dist, found) {
				return true
			}

			for i := range dist {
				if found[i] {
					visited[i] = true
				}
			}
		}
	}

	return false
}

func main() {
	file := flag.String("file", "data.txt", "the file to parse as graphs")
	flag.Parse()

	dat, err := os.ReadFile(*file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n\n")
	k, _ := strconv.Atoi(lines[0])

	res := make([]int, k)
	for i, edgeList := range lines[1:] {
		edges := utils.EdgeListToDirAdjList(strings.Split(edgeList, "\n"))
		weights := utils.WeightsFromEdgeList(strings.Split(edgeList, "\n"))

		if negCycle(edges, weights) {
			res[i] = 1
		} else {
			res[i] = -1
		}
	}

	utils.RosalindPrintArr(res)
}
