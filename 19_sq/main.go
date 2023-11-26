package main

import (
	"flag"
	"fmt"
	"os"
	"rosalind/utils"
	"strconv"
	"strings"
)

// Checks if a simple (non-self-intersecting) square cycle exists in a graph
// defined by an adjacency matrix `mat`.
func squareExists(mat [][]int) bool {
  // We can find a square cycle by identifying four chained nodes, which is
  // equivalent to finding an L shape in the adjacency matrix, and checking if
  // (1) the fourth node is connected to the first node; and (2) if the square
  // cycle doesn't intersect itself
  for i := range mat {
    for j := range mat {
      // For each cell, want to see if there is a 1 to start the square "chain"
      if mat[i][j] == 0 {
        continue
      }
      for r := i+1; r < len(mat); r++ {
        if mat[r][j] == 0 {
          continue
        }
        for c := j+1; c < len(mat); c++ {
          if mat[r][c] == 0 {
            continue
          }
          // Found 4 nodes, check if two the conditions hold:
          // 1. mat[i][c] must be 1 (connected square cycle)
          // 2. mat[i][r] must be 0 and mat[j][c] must be 0 (no intersection)
          if mat[i][c] == 1 && mat[i][r] == 0 && mat[j][c] == 0 {
            return true
          }
        }
      }
    }
  }
  return false
}

func main() {
  file := flag.String("file", "data.txt", "the file to be parsed.")
  flag.Parse()

  dat, err := os.ReadFile(*file)
  if err != nil {
    panic(err)
  }
  s := strings.Split(strings.TrimSpace(string(dat)), "\n\n")
  n, _ := strconv.Atoi(s[0])
  graphs := s[1:]

  res := make([]int, n)
  for i, graph := range graphs {
    if squareExists(utils.EdgeListToUndirAdjMat(strings.Split(graph, "\n"))) {
      res[i] = 1
    } else {
      res[i] = -1
    }
  }
  fmt.Println(res)
}
