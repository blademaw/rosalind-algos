package main

import (
	"flag"
	"fmt"
	"rosalind/utils"
	"strings"
)

func threeWayPartition(xs []int) []int {
  // Choose first item as pivot
  pivot := xs[0]
  i, j := 0, len(xs) - 1
  numEqual := 0
  res := make([]int, len(xs))

  for _, x := range xs[1:] {
    if x < pivot {
      res[i] = x
      i++
    } else if x > pivot {
      res[j] = x
      j--
    } else {
      numEqual++
    }
  }
  for k := i; k <= i + numEqual; k++ {
    res[k] = pivot
  }

  return res
}

func main() {
  filename := flag.String("file", "data.txt", "the file to parse data from.")
  flag.Parse()

  lines, err := utils.ReadLines(*filename)
  if err != nil {
    panic(err)
  }

  xs := utils.StringArrToInts(strings.Split(lines[1], " "))
  fmt.Println(threeWayPartition(xs))
}

