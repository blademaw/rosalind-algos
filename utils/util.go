package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Pretty-prints an array of integers as required by Rosalind.
func RosalindPrintArr(arr []int) {
	fmt.Println(strings.Trim(fmt.Sprint(arr), "[]"))
}

// Converts a list of strings to list of integers.
func StringArrToInts(arr []string) []int {
	res := make([]int, len(arr))
	for i, e := range arr {
		e_int, _ := strconv.Atoi(e)
		res[i] = e_int
	}

	return res
}

// Converts a list of strings in edge list format to an adjacency
// list. Takes `n`, the number of nodes, `arr`, the list of
// strings. Assumes nodes are 1-based indexed.
func EdgesToUndirAdjList(n int, arr []string) [][]int {
	edges := make([][]int, n)
	for _, edge := range arr {
		nodes := strings.Split(edge, " ")
		n1, _ := strconv.Atoi(nodes[0])
		n2, _ := strconv.Atoi(nodes[1])

		n1--
		n2--
		edges[n1] = append(edges[n1], n2)
		edges[n2] = append(edges[n2], n1)
	}

	return edges
}

// Converts a list of strings in edge list format to a directed
// adjacency list.
func EdgesToDirAdjList(n int, arr []string) [][]int {
	edges := make([][]int, n)
	for _, edge := range arr {
		nodes := strings.Split(edge, " ")
		n1, _ := strconv.Atoi(nodes[0])
		n2, _ := strconv.Atoi(nodes[1])

		n1--
		n2--
		edges[n1] = append(edges[n1], n2)
	}

	return edges
}

// Converts a list of strings in edge list format to a directed
// adjacency list.
func EdgeListToDirAdjList(arr []string) [][]int {
	n, _ := strconv.Atoi(strings.Split(arr[0], " ")[0])

	edges := make([][]int, n)
	for _, edge := range arr[1:] {
		nodes := strings.Split(edge, " ")
		n1, _ := strconv.Atoi(nodes[0])
		n2, _ := strconv.Atoi(nodes[1])

		n1--
		n2--
		edges[n1] = append(edges[n1], n2)
	}

	return edges
}

// Converts a list of strings in edge list format to an undirected
// adjacency matrix representation.
func EdgeListToUndirAdjMat(edges []string) [][]int {
	n, _ := strconv.Atoi(strings.Split(edges[0], " ")[0])
	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, n)
	}

	for _, edge := range edges[1:] {
		both := strings.Split(edge, " ")
		n1, _ := strconv.Atoi(both[0])
		n2, _ := strconv.Atoi(both[1])
		n1--
		n2--

		mat[n1][n2] = 1
		mat[n2][n1] = 1
	}

	return mat
}

// Converts a list of strings in edge list format with the number of nodes
// and edges at the start to an adjacency list. Takes `n`, the number of
// nodes, `arr`, the list of strings. Assumes nodes are 1-based indexed.
func EdgeListToUndirAdjList(arr []string) [][]int {
	n, _ := strconv.Atoi(strings.Split(arr[0], " ")[0])

	edges := make([][]int, n)
	for _, edge := range arr[1:] {
		nodes := strings.Split(edge, " ")
		n1, _ := strconv.Atoi(nodes[0])
		n2, _ := strconv.Atoi(nodes[1])

		n1--
		n2--
		edges[n1] = append(edges[n1], n2)
		edges[n2] = append(edges[n2], n1)
	}

	return edges
}

// Converts a list of edges in adjacency list form to weights
func WeightsFromEdgeList(arr []string) [][]int {
	n, _ := strconv.Atoi(strings.Split(arr[0], " ")[0])

	weights := make([][]int, n)
	for i := range weights {
		weights[i] = make([]int, n)
	}

	for _, edge := range arr[1:] {
		nodes := strings.Split(edge, " ")
		n1, _ := strconv.Atoi(nodes[0])
		n2, _ := strconv.Atoi(nodes[1])
		w, _ := strconv.Atoi(nodes[2])

		weights[n1-1][n2-1] = w
	}

	return weights
}

// Loads a file and returns a list of strings.
func ReadLines(filename string) (data []string, err error) {
	dat, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(dat)), "\n"), nil
}
