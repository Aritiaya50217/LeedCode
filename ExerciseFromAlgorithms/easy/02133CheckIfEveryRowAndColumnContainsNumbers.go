package main

import "fmt"

func checkValid(matrix [][]int) bool {
	n := 3
	if len(matrix) < n {
		return true
	}
	list := []int{}
	for _, v := range matrix {
		for i := 0; i < len(v); i++ {
			if v[i] == n {
				list = append(list, v[i])
			}
		}
	}
	if len(list) == n {
		return true
	}
	return false
}

//TODO
func main() {
	num := []int{1, 1, 1}
	matrix := [][]int{}
	matrix = append(matrix, num)
	fmt.Println(matrix)
}
