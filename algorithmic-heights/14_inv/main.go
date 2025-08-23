package main

import (
	"flag"
	"fmt"
	"rosalind/utils"
	"strings"
)

func merge(xs []int, ys []int) (int, []int) {
  n, m := len(xs), len(ys)
  res  := make([]int, n+m)
  inversions := 0

  // Main loop: keep two pointers, add smallest element at each time.
	i, j := 0, 0
	for i < n && j < m {
		if xs[i] <= ys[j] {
			res[i+j] = xs[i]
			i++
		} else {
			res[i+j] = ys[j]
			j++
      inversions += n - i
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

  return inversions, res
}

func mergeSort(xs []int) (int, []int) {
  // Base case
  if len(xs) <= 1 {
    return 0, xs
  }

  // Recursive case
  n := len(xs)
  leftN, left := mergeSort(xs[:n/2])
  rightN, right := mergeSort(xs[n/2:])

  totalN, arr := merge(left, right)
  return (totalN + leftN + rightN), arr
}

func countInversions(xs []int) int {
  num, _ := mergeSort(xs)
  return num
}

func main() {
  filename := flag.String("file", "data.txt", "the data file to parse")
  flag.Parse()

  lines, err := utils.ReadLines(*filename)
  if err != nil {
    panic(err)
  }

  xs := utils.StringArrToInts(strings.Split(lines[1], " "))
  fmt.Println(countInversions(xs))
}
