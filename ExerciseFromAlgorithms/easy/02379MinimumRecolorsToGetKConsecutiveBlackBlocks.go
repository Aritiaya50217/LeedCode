package main

import "fmt"

func minimumRecolors(blocks string, k int) int {
	count := 0
	for i := 0; i < k; i++ {
		if string(blocks[i]) == "W" {
			count++
		}
	}
	return count
}
func main() {
	blocks := "WBWBBBW"
	k := 2
	// count := 0
	for i, _ := range blocks[0:k] {
		fmt.Printf("%v ", string(blocks[i]))
	}

}
