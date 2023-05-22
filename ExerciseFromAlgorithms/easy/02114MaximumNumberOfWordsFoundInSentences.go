package main

import (
	"fmt"
	"strings"
)

func max(array []int, length int) int {
	for i := 0; i < length-1; i++ {
		if array[i] > array[i+1] {
			var temp = array[i]
			array[i] = array[i+1]
			array[i+1] = temp
		}
	}
	var maxValue = array[length-1]
	return maxValue
}
func mostWordsFound(sentences []string) int {
	length := []int{}
	for _, v := range sentences {
		re := strings.Split(v, " ")
		length = append(length, len(re))
	}
	res := max(length, len(length))

	return res
}

func main() {
	sentences := []string{"alice and bob love leetcode", "i think so too",
		"this is great thanks very much"}
	res := mostWordsFound(sentences)
	fmt.Println(res)
}
