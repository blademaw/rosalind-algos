package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Calculates the number of swaps on an insertion sort algorithm.
func countInsertions(ns []int) int {
  num_swaps := 0
  for i := 1; i < len(ns); i++ {
    for k := i; k > 0 && ns[k] < ns[k-1]; k-- {
      ns[k-1], ns[k] = ns[k], ns[k-1]
      num_swaps++
    }
  }
  return num_swaps
}


func main() {
  // Read file in
  dat, err := os.ReadFile("data.txt")
  if err != nil {
    panic(err)
  }
  s := strings.Split(string(dat), "\n")

  // Get n, list of ns
  n, _ := strconv.Atoi(s[0])
  ns_str := strings.Split(s[1], " ")

  ns := make([]int, n)
  for i, v := range ns_str {
    vi, _ := strconv.Atoi(v)
    ns[i]  = vi
  }

  // Do the work
  fmt.Println(countInsertions(ns))
}
