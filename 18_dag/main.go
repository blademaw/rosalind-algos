package main

import (
	"flag"
	"fmt"
	"os"
	"rosalind/utils"
	"strconv"
	"strings"
)

// Conducts a DFS to check if a cycle exists in a graph starting from a node.
func containsCycle(node int, edges [][]int) bool {
  stack, ancestors := utils.Stack{}, utils.Stack{}
  discovered := make([]bool, len(edges))
  
  stack.Push(node)
  ancestors.Push([]int{})

  for !stack.IsEmpty() {
    // Get the current node
    curNode := stack.Pop().(int)
    curAncs := ancestors.Pop().([]int)
    newAncs := append(curAncs, curNode)

    // Set it to discovered
    discovered[curNode] = true

    // For each child, add if not discovered, and if the child is in the
    // ancestor list, we have a cycle.
    for _, child := range edges[curNode] {
      for _, parent := range newAncs {
        if child == parent {
          return true
        }
      }
      if !discovered[child] {
        stack.Push(child)
        ancestors.Push(newAncs)
      }
    }
  }
  return false
}


// Converts a list of strings in edge list format to a directed adjacency list
// of both in-edges and out-edges
func edgesToInOutEdges(n int, arr []string) ([][]int, [][]int) {
  inEdges, outEdges := make([][]int, n), make([][]int, n)
  for _, edge := range arr {
    nodes := strings.Split(edge, " ")
    n1, _ := strconv.Atoi(nodes[0])
    n2, _ := strconv.Atoi(nodes[1])

    n1--; n2--
    inEdges[n1] = append(inEdges[n1], n2)
    outEdges[n2] = append(inEdges[n2], n1)
  }
  
  return inEdges, outEdges
}

// Detects whether a graph (represented by an edge list array) is acyclic or
// not.
func isDAG(outEdges [][]int, inEdges [][]int) bool {
  // DFS from each root (node without any in-edges -- if doesn't exist, have
  // cycle) until we hit a node with a back-edge (has a child that is a
  // parent).
  for i, edges := range inEdges {
    // Node without any in-edges is a root
    if len(edges) == 0 {
      if containsCycle(i, outEdges) {
        return false
      }
    }
  }
  return true
}

func main() {
  filename := flag.String("file", "data.txt", "the filename to parse.")
  flag.Parse()

  dat, err := os.ReadFile(*filename)
  if err != nil {
    panic(err)
  }
  s := strings.Split(strings.TrimSpace(string(dat)), "\n\n")
  nDags, _ := strconv.Atoi(s[0])
  dagStrs := s[1:]

  // Turn into directed edge lists
  outDags, inDags := make([][][]int, nDags), make([][][]int, nDags)
  for i := range outDags {
    n, _ := strconv.Atoi(strings.Split(strings.Split(dagStrs[i], "\n")[0], " ")[0])
    outDags[i], inDags[i] = edgesToInOutEdges(n, strings.Split(dagStrs[i], "\n")[1:])
  }

  // For each graph, compute if DAG or not
  res := make([]int, nDags)
  for i := range inDags {
    switch isDAG(outDags[i], inDags[i]) {
      case true:
        res[i] = 1
      case false:
        res[i] = -1
    }
  }

  fmt.Println(res)
}
