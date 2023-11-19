package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findMaxNodes(edges []string) int {
  maxNodes := 0 // NOTE: starting to adopt Go naming standards...
  for _, edge := range edges {
    nodes := strings.Split(edge, " ")
    n1, _ := strconv.Atoi(nodes[0])
    n2, _ := strconv.Atoi(nodes[1])

    if n1 > maxNodes {
      maxNodes = n1
    } else if n2 > maxNodes {
      maxNodes = n2
    }
  }

  return maxNodes
}

func main() {
  dat, err := os.ReadFile("data.txt")
  if err != nil {
    panic(err)
  }

  // Iterating once to find num of nodes
  s := strings.Split(strings.TrimSpace(string(dat)), "\n")
  n, _ := strconv.Atoi(strings.Split(s[0], " ")[0])
  edges := make([][]int, n)

  // Creating the adjacency array
  for _, edge := range s[1:] {
    nodes := strings.Split(edge, " ")
    n1, _ := strconv.Atoi(nodes[0])
    n2, _ := strconv.Atoi(nodes[1])
    n1--
    n2--

    edges[n1] = append(edges[n1], n2)
    edges[n2] = append(edges[n2], n1)
  }

  // Outputting the answer
  for i := range edges {
    total := 0
    for _, neighbor := range edges[i] {
      total += len(edges[neighbor])
    }
    fmt.Printf("%d ", total)
  }
  fmt.Printf("\n")
  
}
