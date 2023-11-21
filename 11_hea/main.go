package main

import (
	"fmt"
	"os"
	"rosalind/utils"
	"strings"
)

func bubbleUp(heap []int, x int, i int) {
  heap[i] = x
  j := (i-1)/2

  for j >= 0 {
    if heap[j] < x {
      heap[i], heap[j] = heap[j], heap[i]
      i = j
      j = (j-1)/2
    } else {
      break
    }
  }
}

func makeHeapFromArr(arr []int) []int {
  heap := make([]int, len(arr))
  heap[0] = arr[0]

  for i := 1; i < len(heap); i++ {
    bubbleUp(heap, arr[i], i)
  }

  return heap
}

func main() {
  dat, err := os.ReadFile("data.txt")
  if err != nil {
    panic(err)
  }

  lines := strings.Split(string(dat), "\n")
  arr := utils.StringArrToInts(strings.Split(lines[1], " "))

  fmt.Println(makeHeapFromArr(arr))
}
