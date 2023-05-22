package main

import "fmt"

func construct2DArray(original []int, m int, n int) [][]int {
	twoDimension := [][]int{}
	if len(original) >= 4 && len(original)%2 == 0 {
		twoDimension = append(twoDimension, original[:m], original[n:])
	} else if len(original)%3 == 0 {
		twoDimension = append(twoDimension, original[n:])
	}
	return twoDimension
}

// TODO
func main() {
	original := []int{1, 2, 3}
	m, n := 1, 3
	twoDimension := [][]int{}
	if len(original) >= 4 && len(original)%2 == 0 {
		twoDimension = append(twoDimension, original[:m], original[n:])
	} else if len(original)%3 == 0 {
		twoDimension = append(twoDimension, original[n:])
	}
	fmt.Println(twoDimension, n, m)

}
