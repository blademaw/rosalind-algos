package main

import (
	"flag"
	"fmt"
	"os"
	"rosalind/utils"
	"strconv"
	"strings"
)

func dijkstra(s int, edges [][]int, weights [][]int) []int {
	// Implements Dijkstra's algorithm using a min-heap-based priority queue that
	// starts at node `s`. Returns distance array where index `i` represents the
	// distance from `s` to node `i`.

	n := len(edges)

	distances := make([]int, n)
	for i := range distances {
		distances[i] = -1
	}
	distances[s] = 0

	pQueue := utils.NewPriorityQueue()
	pQueue.Insert(s, 0)

	for !pQueue.IsEmpty() {
		u, key, ok := pQueue.PopMin()
		if !ok {
			panic("Error in retrieving minimum from priority queue.")
		}

		if distances[u] == key {
			for _, v := range edges[u] {
				if distances[v] == -1 || distances[v] > distances[u]+weights[u][v] {
					distances[v] = distances[u] + weights[u][v]
					pQueue.Insert(v, distances[v])
				}
			}
		}
	}

	return distances
}

func shortestCycle(u int, v int, edges [][]int, weights [][]int) int {
	// Finds the shortest cycle through the edge (u, v) by doing Dijkstra from v
	// to find u.
	if dists := dijkstra(v, edges, weights); dists[u] == -1 {
		return -1
	} else {
		return weights[u][v] + dists[u]
	}
}

func main() {
	file := flag.String("file", "data.txt", "the file with graphs to parse.")
	flag.Parse()

	dat, err := os.ReadFile(*file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n\n")
	k, _ := strconv.Atoi(lines[0])

	res := make([]int, k)
	for i, graph := range lines[1:] {
		firstEdge := strings.Split(strings.Split(graph, "\n")[1], " ")
		u, _ := strconv.Atoi(firstEdge[0])
		v, _ := strconv.Atoi(firstEdge[1])
		u--
		v--

		edges := utils.EdgeListToDirAdjList(strings.Split(graph, "\n"))
		weights := utils.WeightsFromEdgeList(strings.Split(graph, "\n"))

		res[i] = shortestCycle(u, v, edges, weights)
	}

	fmt.Println(res)
}
