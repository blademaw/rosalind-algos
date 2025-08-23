package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
  // Load the file
  dat, err := os.ReadFile("data.txt")
  if err != nil {
    panic(err)
  }

  // Initialize params
  s := strings.Split(strings.TrimSpace(string(dat)), "\n")
  n, _ := strconv.Atoi(strings.Split(s[0], " ")[0])

  // Make degree count list, iterate
  degs := make([]int, n)
  for _, e := range s[1:] {
    e_vals := strings.Split(e, " ")
    u, _ := strconv.Atoi(e_vals[0])
    v, _ := strconv.Atoi(e_vals[1])

    degs[u-1]++
    degs[v-1]++
  }

  fmt.Println(degs)

}
