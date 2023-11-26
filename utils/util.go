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

// Converts a list of strings in edge list format to an undirected
// adjacency matrix representation.
func EdgeListToUndirAdjMat(edges []string) [][]int {
  n, _ := strconv.Atoi(strings.Split(edges[0], " ")[0])
  mat := make([][]int, n)
  for i := range mat {
    mat[i] = make([]int, n)
  }

  for _, edge := range edges[1:] {
    both  := strings.Split(edge, " ")
    n1, _ := strconv.Atoi(both[0])
    n2, _ := strconv.Atoi(both[1])
    n1--; n2--

    mat[n1][n2] = 1
    mat[n2][n1] = 1
  }

  return mat
}

// Loads a file and returns a list of strings.
func ReadLines(filename string) (data []string, err error) {
  dat, err := os.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  return strings.Split(strings.TrimSpace(string(dat)), "\n"), nil
}
