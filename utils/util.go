package utils

import (
	"os"
	"strconv"
	"strings"
)

// Converts a list of strings to list of integers.
func StringArrToInts(arr []string) []int {
	res := make([]int, len(arr))
	for i, e := range arr {
		e_int, _ := strconv.Atoi(e)
		res[i] = e_int
	}

	return res
}

// Converts a list of strings in edge list format to an adjacency
// list. Takes `n`, the number of nodes, `arr`, the list of
// strings. Assumes nodes are 1-based indexed.
func EdgesToUndirAdjList(n int, arr []string) [][]int {
  edges := make([][]int, n)
  for _, edge := range arr {
    nodes := strings.Split(edge, " ")
    n1, _ := strconv.Atoi(nodes[0])
    n2, _ := strconv.Atoi(nodes[1])

    n1--; n2--
    edges[n1] = append(edges[n1], n2)
    edges[n2] = append(edges[n2], n1)
  }
  
  return edges
}

// Converts a list of strings in edge list format to a directed
// adjacency list.
func EdgesToDirAdjList(n int, arr []string) [][]int {
  edges := make([][]int, n)
  for _, edge := range arr {
    nodes := strings.Split(edge, " ")
    n1, _ := strconv.Atoi(nodes[0])
    n2, _ := strconv.Atoi(nodes[1])

    n1--; n2--
    edges[n1] = append(edges[n1], n2)
  }
  
  return edges
}

// Loads a file and returns a list of strings.
func ReadLines(filename string) (data []string, err error) {
  dat, err := os.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  return strings.Split(strings.TrimSpace(string(dat)), "\n"), nil
}
