package main

import (
	"flag"
	"fmt"
	"rosalind/utils"
)

func bellmanFord(s int, edges [][]int, weights[][]int) []int {
  distance := make([]int, len(edges))
  for i := range distance {
    distance[i] = -1
  }
  distance[s] = 0

  for i := 0; i < len(edges); i++ {
    for u, uEdges := range edges {
      if distance[u] == -1 {
        continue
      }

      for _, v := range uEdges {
        if distance[v] == -1 || distance[v] > distance[u] + weights[u][v] {
          distance[v] = distance[u] + weights[u][v]
        }
      }
    }
  }

  return distance
}

func main() {
  file := flag.String("file", "data.txt", "the file to parse as a graph.")
  flag.Parse()

  lines, err := utils.ReadLines(*file)
  if err != nil {
    panic(err)
  }

  edges := utils.EdgeListToDirAdjList(lines)
  weights := utils.WeightsFromEdgeList(lines)

  for _, d := range bellmanFord(0, edges, weights) {
    if d == -1 {
      fmt.Print("x ")
    } else {
      fmt.Printf("%d ", d)
    }
  }
  fmt.Printf("\n")
}
