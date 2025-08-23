package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"rosalind/utils"
	"strconv"
	"strings"
	"time"
)

var source  = rand.NewSource(time.Now().UnixNano())
var randGen = rand.New(source)

func threeWayPartition(xs []int) ([]int, int, int) {
  // Choose random item in xs as pivot
  pivotInd := randGen.Intn(len(xs))
  xs[pivotInd], xs[0] = xs[0], xs[pivotInd]
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

  return res, i, len(xs) - i - numEqual - 1
}

func selection(s []int, k int) int {
  // Retrieves the `k`-th smallest item from the array `s`. Starts by
  // partitioning the array into three sections, and recursively takes the
  // `k`-th smallest element from the appropriate list.

  // Since the three-way partition is O(n), and recursing on the appropriate
  // sublist is done on either the left or right sublist, if sLeft and sRight
  // are the left and right sublists, we reduce |S| to max(|sLeft|, |sRight|)
  // which is approximately |S|/2. However, this requires picking the median as
  // the partition pivot. Instead, we choose an element randomly.
  
  res, left, right := threeWayPartition(s)
  numEqual := len(res) - left - right

  if k <= left {
    return selection(res[:left], k)
  } else if k > left && k <= left + numEqual {
    return res[left]
  } else {
    return selection(res[len(res)-right:], k - left - numEqual)
  }

}

func main() {
  file := flag.String("file", "data.txt", "the filename to parse as data.")
  flag.Parse()

  dat, err := os.ReadFile(*file)
  if err != nil {
    panic(err)
  }
  lines := strings.Split(strings.TrimSpace(string(dat)), "\n")
  
  ns := utils.StringArrToInts(strings.Split(lines[1], " "))
  k,_:= strconv.Atoi(lines[2])

  fmt.Println(selection(ns, k))
}
