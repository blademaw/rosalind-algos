package main

import (
	"flag"
	"os"
	"rosalind/utils"
	"strconv"
	"strings"
)

func dfs(u int, edges [][]int, visited []bool) {
	visited[u] = true

	for _, u := range edges[u] {
		if !visited[u] {
			dfs(u, edges, visited)
		}
	}
}

func generalSink(edges [][]int) int {
	// This is the problem of finding a "mother vertex"; it can be solved by
	// using the idea from Kosaraju's algorithm for strongly-connected
	// components, where we note that if we iteratively DFS until all vertices
	// are discovered, the last origin node DFS-ed from `v` is a candidate mother
	// vertex. We just need to BFS/DFS from `v` again to ensure we can reach
	// every other node => O(V + E) + O(V + E) = O(V + E)

	visited := make([]bool, len(edges))
	u := 0

	for v := range edges {
		if !visited[v] {
			dfs(v, edges, visited)
			u = v
		}
	}

	visited = make([]bool, len(edges))
	dfs(u, edges, visited)

	for _, visit := range visited {
		if !visit {
			return -1
		}
	}

	return u + 1
}

func main() {
	file := flag.String("file", "data.txt", "the file to parse as a graph.")
	flag.Parse()

	dat, err := os.ReadFile(*file)
	if err != nil {
		panic(err)
	}

	s := strings.Split(strings.TrimSpace(string(dat)), "\n\n")
	k, _ := strconv.Atoi(s[0])

	res := make([]int, k)
	for i, graph := range s[1:] {
		edges := utils.EdgeListToDirAdjList(strings.Split(graph, "\n"))
		res[i] = generalSink(edges)
	}

	utils.RosalindPrintArr(res)
}
