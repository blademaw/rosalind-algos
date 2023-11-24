package main

import (
	"flag"
	"fmt"
	"rosalind/utils"
	"strings"
)

func minChild(heap []int, i int, capacity int) int {
  c1, c2 := i*2 + 1, i*2 + 2
  if c1 >= capacity {
    return -1
  } else if c2 >= capacity {
    return c1
  }

  if heap[c1] > heap[c2] {
    return c1
  }
  return c2
}

func siftDown(heap []int, x int, i int, capacity int) {
  c := minChild(heap, i, capacity)
  for c != -1 && heap[c] > x {
    heap[i] = heap[c]
    i = c
    c = minChild(heap, i, capacity)
  }
  heap[i] = x
}

func deleteMin(heap []int, capacity int) int {
  x := heap[0]
  heap[0], heap[capacity - 1] = heap[capacity - 1], 0
  siftDown(heap, heap[0], 0, capacity - 1)
  return x
}

func heapSort(heap []int) []int {
  capacity := len(heap)
  res := make([]int, capacity)

  for i := capacity-1; i >= 0; i-- {
    res[i] = deleteMin(heap, capacity)
    capacity--
  }

  return res
}

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
  filename := flag.String("filename", "data.txt", "file to parse as data.")
  flag.Parse()

  lines, err := utils.ReadLines(*filename)
  if err != nil {
    panic(err)
  }

  arr := utils.StringArrToInts(strings.Split(lines[1], " "))
  fmt.Println(heapSort(makeHeapFromArr(arr)))
}

