package main

import (
	"fmt"
	"rosalind/utils"
	"strconv"
	"strings"
)

type NodeDepth struct {
  Node  int
  Depth int
}

func shortestPaths(edges [][]int) []int {
  // BFS to find shortest path from vertex 1
  shortestLens := make([]int, len(edges))
  for i := range shortestLens {
    shortestLens[i] = -1
  }

  queue := utils.Queue{}
  queue.Enqueue(NodeDepth{Node: 0, Depth: 0}) // From the first node
  discovered := make([]bool, len(edges))

  for !queue.IsEmpty() {
    pair, ok := queue.Dequeue().(NodeDepth)
    node  := pair.Node
    depth := pair.Depth

    // If seen this node before (or error), move on
    if discovered[node] || !ok {
      continue
    }

    // Otherwise, add children and new depths
    shortestLens[node] = depth
    discovered[node] = true

    for _, child := range edges[node] {
      if !discovered[child] {
        queue.Enqueue(NodeDepth{Node: child, Depth: depth + 1})
      }
    }
  }

  return shortestLens
}

func main() {
  lines, err := utils.ReadLines("data.txt")
  if err != nil {
    panic(err)
  }

  n, _ := strconv.Atoi(strings.Split(lines[0], " ")[0])
  edges := utils.EdgesToDirAdjList(n, lines[1:])

  fmt.Println(shortestPaths(edges))
}
