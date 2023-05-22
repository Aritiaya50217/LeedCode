package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func plusOne(digits []int) []int {
	for _, v := range digits[len(digits)-1:] {
		v++
		digits = append(digits[:len(digits)-1], v)
		if v > 9 {
			return digits
		}
	}
	return digits
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	list := []int{}
	for input.Scan() {
		inputStr := input.Text()
		number, _ := strconv.Atoi(inputStr)
		if inputStr != "" {
			list = append(list, number)
		} else {
			break
		}
	}
	r := plusOne(list)
	fmt.Println(r)
}
