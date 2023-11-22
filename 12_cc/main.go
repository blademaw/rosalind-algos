package main

import (
	"fmt"
	"os"
	"rosalind/utils"
	"strconv"
	"strings"
)

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

func numConnectedComponents(n int, edges [][]int) int {
	// Initialize all to discovered
	discovered := make([]bool, n)
	components := 0

	// For each node, DFS to identify component number
	for i := 0; i < n; i++ {
		// If this node already has an identified component, skip
		if discovered[i] {
			continue
		}

		dfs(i, edges, discovered)

		components++
	}
	return components
}

func main() {
	dat, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	s := strings.Split(strings.TrimSpace(string(dat)), "\n")
	n, _ := strconv.Atoi(strings.Split(s[0], " ")[0])

	adjArr := utils.EdgeListToAdjacencyList(n, s[1:])
	fmt.Println("Number of components is", numConnectedComponents(n, adjArr))
}
