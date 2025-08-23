package main

import (
  "flag"
  "fmt"
  "rosalind/utils"
  "strings"
)

func sumThreeInds(xs []int) {
  // for each index i, j, want to find some k such that A[i]+A[j] = -A[k],
  // so their sum is 0. Can do this by hashing num -> index, and checking
  // if -(A[i]+A[j]) exists in hash map for every i, j > i
  n := len(xs)
  m := make(map[int]int)
  for i, elem := range xs {
    m[elem] = i
  }

  for i := 0; i < n; i++ {
    for j := i; j < n; j++ {
      k, ok := m[-(xs[i] + xs[j])]
      if !ok {
        continue
      }
      fmt.Println(i + 1, j + 1, k + 1)
      return
    }
  }
  fmt.Println("-1")
}

func main() {
  filename := flag.String("file", "data.txt", "the data file to parse.")
  flag.Parse()

  lines, err := utils.ReadLines(*filename)
  if err != nil {
    panic(err)
  }

  for _, arr := range lines[1:] {
    xs := utils.StringArrToInts(strings.Split(arr, " "))
    sumThreeInds(xs)
  }
}
