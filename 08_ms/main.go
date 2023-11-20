package main

import (
	"fmt"
	"os"
	"rosalind/utils"
	"strings"
)

func merge(xs []int, ys []int) []int {
  n, m := len(xs), len(ys)
  res  := make([]int, n+m)

  // Main loop: keep two pointers, add smallest element at each time.
	i, j := 0, 0
	for i < n && j < m {
		if xs[i] < ys[j] {
			res[i+j] = xs[i]
			i++
		} else {
			res[i+j] = ys[j]
			j++
		}
	}

	// If we haven't added everything from one array, add the rest
	if i < n {
		for i+j < n+m {
			res[i+j] = xs[i]
			i++
		}
	} else if j < m {
		for i+j < n+m {
			res[i+j] = ys[j]
			j++
		}
	}

  return res
}

func mergeSort(xs []int) []int {
  // Base case
  if len(xs) <= 1 {
    return xs
  }

  // Recursive case
  n := len(xs)
  left, right := mergeSort(xs[:n/2]), mergeSort(xs[n/2:])

  return merge(left, right)
}

func main() {
  dat, err := os.ReadFile("data.txt")
  if err != nil {
    panic(err)
  }

  s := strings.Split(string(dat), "\n")
  ns := utils.StringArrToInts(strings.Split(s[1], " "))

  fmt.Println(mergeSort(ns))
}
