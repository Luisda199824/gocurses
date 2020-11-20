package main

import "fmt"

func main() {
	var solutions [][]int
	var tempSolution = []int{1, 2}
	solutions = append(solutions, tempSolution)
	tempSolution = []int{3, 4}
	solutions = append(solutions, tempSolution)

	fmt.Println(solutions)
	fmt.Println(solutions[:1])
}
