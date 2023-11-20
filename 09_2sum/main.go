package main

import (
	"fmt"
	"os"
	"rosalind/utils"
	"strings"
)

func printTwoSumIndices(xs []int) {
  m := make(map[int]int)
  for i := 0; i < len(xs); i++ {
    // Check if negated term exists
    ind, ok := m[-xs[i]]
    if ok {
      fmt.Println(ind+1, i+1)
      return
    }

    // If doesn't exist, add current index
    m[xs[i]] = i
  }

  // If no indices, return -1
  fmt.Println("-1")
}

func main() {
  dat, err := os.ReadFile("data.txt")
  if err != nil {
    panic(err)
  }

  s := strings.Split(strings.TrimSpace(string(dat)), "\n")
  for _, strArr := range s[1:] {
    arr := utils.StringArrToInts(strings.Split(strArr, " "))
    printTwoSumIndices(arr)
  }
}
