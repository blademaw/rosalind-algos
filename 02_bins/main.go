package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fibonnaci
func fibonacci_numbers(n int) int {
  fibs := make([]int, n+1)
  fibs[0] = 0
  fibs[1] = 1

  for i := 2; i <= n; i++ {
    fibs[i] = fibs[i-1] + fibs[i-2]
  }
  return fibs[n]
}


// Single binary search. Takes an integer as input with a sorted list of
// integers, and returns the position in the array (with one-based indexing) if
// it can be found, or -1 if not.
func binary_search(n int, ns []int) int {
  low, i, high := 0, len(ns) / 2, len(ns)

  for low < high {
    i = (low + high) / 2
    k := ns[i]

    if n == k {
      return i + 1
    } else if n < k {
      high = i
    } else {
      low = i + 1
    }
  }

  return -1
}


func main() {
  // NOTE: this error checking is verbose and extensive, I know I can consume
  // the value with var, _ := ...; but I am learning Go.

  // open the file
  dat, err := os.ReadFile("data.txt")
  if err != nil {
    panic(err)
  }
  dat_str := strings.Split(string(dat), "\n")

  // grab n, m
  n, err := strconv.Atoi(dat_str[0])
  if err != nil {
    fmt.Println("Converting n failed:", err)
  }

  m, err := strconv.Atoi(dat_str[1])
  if err != nil {
    fmt.Println("Converting m failed:", err)
  }

  // parse list of strs
  ns_str := strings.Split(dat_str[2], " ")
  ns := make([]int, n)
  for i, val := range ns_str {
    new_val, _ := strconv.Atoi(val)
    ns[i] = new_val
  }

  ms_str := strings.Split(dat_str[3], " ")
  ms := make([]int, m)
  for i, val := range ms_str {
    new_val, _ := strconv.Atoi(val)
    ms[i] = new_val
  }

  // loop to print solution
  for _, val := range ms {
    fmt.Print(binary_search(val, ns), " ")
  }
  fmt.Println()
}
