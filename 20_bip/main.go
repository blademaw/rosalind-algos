package main

import (
	"flag"
	"fmt"
	"os"
	"rosalind/utils"
	"strconv"
	"strings"
)

func isBipartite(edges [][]int) bool {
  // Implements the two-coloring algorithm to test if a graph in adjacency
  // array format is bipartite or not. A graph is bipartite if it can be
  // colored using two colors.

  // To solve this problem, we start by assigning a random node a color and
  // BFSing from it (assume graph is connected). At each new node discovered,
  // we retrieve all adjacent nodes that are discovered and have a color. If
  // the colors are all the same, we assign the current node the opposite
  // color. Otherwise, if any colors disagree, the graph is not bipartite.
  discovered := make([]bool, len(edges))
  colors := make([]int, len(edges)) // 0 is unassigned, 1/2 are colors

  queue := utils.Queue{}
  queue.Enqueue(0)

  for !queue.IsEmpty() {
    node, ok := queue.Dequeue().(int)
    if discovered[node] || !ok {
      continue
    }

    discovered[node] = true

    // Find all adjacent colors
    var curColors []int
    for _, child := range edges[node] {
      if discovered[child] {
        curColors = append(curColors, colors[child])
      }
    }

    // Ensure they are the same
    var startColor int
    if len(curColors) == 0 {
      startColor = 1
    } else {
      startColor = curColors[0]
      for _, c := range curColors[1:] {
        if c != startColor {
          return false
        }
      }
    }

    // Assign the opposite color, add nodes to queue
    if startColor == 1 {
      colors[node] = 2
    } else {
      colors[node] = 1
    }

    for _, child := range edges[node] {
      if !discovered[child] {
        queue.Enqueue(child)
      }
    }
  }

  return true
}

func main() {
  filename := flag.String("file", "data.txt", "the file to parse as data.")
  flag.Parse()

  dat, err := os.ReadFile(*filename)
  if err != nil {
    panic(err)
  }

  graphs := strings.Split(strings.TrimSpace(string(dat)), "\n\n")
  n, _ := strconv.Atoi(graphs[0])
  res := make([]int, n)

  for i, g := range graphs[1:] {
    if isBipartite(utils.EdgeListToUndirAdjList(strings.Split(g, "\n"))) {
      res[i] = 1
    } else {
      res[i] = -1
    }
  }

  fmt.Println(res)
}
