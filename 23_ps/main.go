package main

import (
	"flag"
	"fmt"
	"rosalind/utils"
	"strconv"
	"strings"
)

func partialSort(ns []int, k int) []int {
  pQueue := utils.NewPriorityQueue()
  for _, n := range ns {
    pQueue.Insert(n, n)
  }

  res := make([]int, k)
  for i := 0; i < k; i++ {
    num, _, ok := pQueue.PopMin()
    if !ok {
      panic("Could not pop min from heap.")
    }

    res[i] = num
  }
  return res
}

func main() {
  file := flag.String("file", "data.txt", "the file to parse as an array.")
  flag.Parse()

  lines, err := utils.ReadLines(*file)
  if err != nil {
    panic(err)
  }

  k, _ := strconv.Atoi(lines[2])
  xs := utils.StringArrToInts(strings.Split(lines[1], " "))

  fmt.Println(partialSort(xs, k))
}
