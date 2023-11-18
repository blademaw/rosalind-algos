package main

import (
	"fmt"
	"os"
	"rosalind/utils"
	"strconv"
	"strings"
)

func main() {
	// Read the file
	dat, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	s := strings.Split(string(dat), "\n")

	n, _ := strconv.Atoi(s[0])
	m, _ := strconv.Atoi(s[2])

	ns := utils.StringArrToInts(strings.Split(s[1], " "))
	ms := utils.StringArrToInts(strings.Split(s[3], " "))

	res := make([]int, n+m)

	// Main loop: keep two pointers, add smallest element at each time.
	i, j := 0, 0
	for i < n && j < m {
		if ns[i] < ms[j] {
			res[i+j] = ns[i]
			i++
		} else {
			res[i+j] = ms[j]
			j++
		}
	}

	// If we haven't added everything from one array, add the rest
	if i < n {
		for i+j < n+m {
			res[i+j] = ns[i]
			i++
		}
	} else if j < m {
		for i+j < n+m {
			res[i+j] = ms[j]
			j++
		}
	}

	fmt.Println(res)
}
