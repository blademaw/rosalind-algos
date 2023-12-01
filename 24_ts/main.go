package main

import (
	"flag"
	"fmt"
	"rosalind/utils"
)

// Implements a topological sort for DAGs. Works by DFSing and recording
// post-order values according to a clock (how long nodes have been on stack)
// and returning the nodes in descending order (sinks have smallest post
// values, sources have largest).
func topologicalSort(edges [][]int) []int {
  n := len(edges)

  // DFS, insert in order visited
  discovered := make([]bool, n)
  res, i := make([]int, n), n - 1

  var dfs func(u int)
  dfs = func(u int)  {
    discovered[u] = true

    for _, v := range edges[u] {
      if !discovered[v] {
        dfs(v)
      }
    }
    res[i] = u + 1 // Rosalind formatting
    i--
  }

  for u := range edges {
    if !discovered[u] {
      dfs(u)
    }
  }

  return res
}

func main() {
  file := flag.String("file", "data.txt", "the file to be parsed as a DAG.")
  flag.Parse()

  lines, err := utils.ReadLines(*file)
  if err != nil {
    panic(err)
  }

  edges := utils.EdgeListToDirAdjList(lines)

  fmt.Println(topologicalSort(edges))
}
