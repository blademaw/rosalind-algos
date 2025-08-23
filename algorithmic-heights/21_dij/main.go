package main

import (
	"flag"
	"fmt"
	"os"
	"rosalind/utils"
	"strconv"
	"strings"
)

func dijkstra(s int, edges [][]int, weights [][]int) []int {
  // Implements Dijkstra's algorithm using a min-heap-based priority queue that
  // starts at node `s`. Returns distance array where index `i` represents the
  // distance from `s` to node `i`.

  n := len(edges)

  distances := make([]int, n)
  for i := range distances {
    distances[i] = -1
  }
  distances[s] = 0

  pQueue := utils.NewPriorityQueue()
  pQueue.Insert(s, 0)

  for !pQueue.IsEmpty() {
    u, key, ok := pQueue.PopMin()
    if !ok {
      panic("Error in retrieving minimum from priority queue.")
    }

    if distances[u] == key {
      for _, v := range edges[u] {
        if distances[v] == -1 || distances[v] > distances[u] + weights[u][v] {
          distances[v] = distances[u] + weights[u][v]
          pQueue.Insert(v, distances[v])
        }
      }
    }
  }

  return distances
}

func main() {
  filename := flag.String("file", "data.txt", "the file to be parsed for data.")
  flag.Parse()

  dat, err := os.ReadFile(*filename)
  if err != nil {
    panic(err)
  }

  lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
  n, _ := strconv.Atoi(strings.Split(lines[0], " ")[0])
  
  adjMat := make([][]int, n)
  for i := range adjMat {
    adjMat[i] = make([]int, n)
  }

  edges := make([][]int, n)
  for _, edge := range lines[1:] {
    nodes := strings.Split(edge, " ")
    n1, _ := strconv.Atoi(nodes[0])
    n2, _ := strconv.Atoi(nodes[1])
    w , _ := strconv.Atoi(nodes[2])

    n1--; n2--
    edges[n1] = append(edges[n1], n2)
    adjMat[n1][n2] = w
  }

  fmt.Println(dijkstra(0, edges, adjMat))
}
