package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
  "rosalind/utils"
)

// Finds the majority element in a list. Returns -1 if doesn't exist.
func majorityElement(arr []int) int {
  counts, half_size := make(map[int]int), len(arr)/2
  for _, elem := range arr {
    counts[elem]++
    if counts[elem] > half_size {
      return elem
    }
  }

  return -1
}

func main() {
  // Read the file
  dat, err := os.ReadFile("data.txt")
  if err != nil {
    panic(err)
  }

  // Get parameters
  s := strings.Split(strings.TrimSpace(string(dat)), "\n")
  k, _ := strconv.Atoi(strings.Split(s[0], " ")[0])
  k_arr:= make([]int, k)

  // For each arr, find if majority element and set
  for i, arr := range s[1:] {
    i_arr := utils.StringArrToInts(strings.Split(arr, " "))
    k_arr[i] = majorityElement(i_arr)
  }

  fmt.Println(k_arr)
}
