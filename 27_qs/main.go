package main

import (
	"flag"
	"fmt"
	"rosalind/utils"
	"strings"
)

func partition(xs []int, low int, high int) (int, int) {
	newLow, newHigh := low, high
	i := low
	pivot := xs[low]

	for i <= newHigh {
		if xs[i] < pivot {
			xs[newLow], xs[i] = xs[i], xs[newLow]
			i++
			newLow++
		} else if xs[i] > pivot {
			xs[newHigh], xs[i] = xs[i], xs[newHigh]
			newHigh--
		} else {
			i++
		}
	}

	return newLow, newHigh
}

func quicksort(arr []int, low int, high int) {
	// Quicksort an array in-place.

	if low < high {
		left, right := partition(arr, low, high)

		quicksort(arr, low, left-1)
		quicksort(arr, right+1, high)
	}
}

func main() {
	file := flag.String("file", "data.txt", "the data to parse as a list.")
	flag.Parse()

	lines, err := utils.ReadLines(*file)
	if err != nil {
		panic(err)
	}

	arr := utils.StringArrToInts(strings.Split(lines[1], " "))
	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
