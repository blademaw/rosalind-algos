package main

import "fmt"

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

func main() {
  fmt.Println(fibonacci_numbers(24))
}
