package main

import "fmt"

// ref :https://www.tutorialspoint.com/write-a-golang-program-to-find-duplicate-elements-in-a-given-array
func duplicateInArray(arr []int) int {
	check := make(map[int]bool, 0)
	for i := 0; i < len(arr); i++ {
		if check[arr[i]] == true {
			return arr[i]
		} else {
			check[arr[i]] = true
		}
	}
	return -1
}
// ref : https://www.tutorialspoint.com/go/go_multi_dimensional_arrays.htm
func findCenter(edges [][]int) int {
	list := []int{}
	for i := 0; i < len(edges); i++ {
		for j := 0; j < 2; j++ {
			number := edges[i][j]
			list = append(list, number)
		}
	}
	result := duplicateInArray(list)
	return result
}
func main() {
	num := []int{1, 2}
	num2 := []int{2, 3}
	num3 := []int{4, 2}
	edges := [][]int{}
	edges = append(edges, num, num2, num3)

	result := findCenter(edges)
	fmt.Println(result)
}
