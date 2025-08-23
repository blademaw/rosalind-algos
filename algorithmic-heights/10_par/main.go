package main

import (
	"fmt"
	"os"
	"rosalind/utils"
	"strings"
)

// Partition an array of integers, taking the first element as the pivot.
// Assumes n >= 3.
func partition(xs []int) []int {
  // Choose first item as pivot
  pivot := xs[0]
  i, j := 0, len(xs) - 1
  res := make([]int, len(xs))

  for _, x := range xs[1:] {
    if x <= pivot {
      res[i] = x
      i++
    } else {
      res[j] = x
      j--
    }
  }
  res[i] = pivot

  return res
}

func main() {
  dat, err := os.ReadFile("data.txt")
  if err != nil {
    panic(err)
  }

  s := strings.Split(string(dat), "\n")
  ns := utils.StringArrToInts(strings.Split(s[1], " "))

  fmt.Println(partition(ns))
}
