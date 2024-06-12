package main

import (
	"fmt"
	"strconv"
	"strings"
)

func plusOne(digits []int) []int {
	last := digits[len(digits)-1:]
	num := last[0] + 1
	digits = digits[:len(digits)-1]
	digits = append(digits, num)

	if last[0] > 9 {
		first := digits[:len(digits)-1]
		numberStr := strconv.Itoa(last[0])
		sp := strings.Split(numberStr, "")
		for _, v := range sp {
			s, _ := strconv.Atoi(v)
			first = append(first, s)
		}
		return first
	}
	return digits
}

// TODO
func main() {
	digits := []int{9, 9}

	// last := digits[len(digits)-1:]
	// num := last[0] + 1
	// digits = digits[:len(digits)-1]
	// digits = append(digits, num)

	// if last[0] > 9 {
	// 	first := digits[:len(digits)-1]
	// 	numberStr := strconv.Itoa(last[0])
	// 	sp := strings.Split(numberStr, "")
	// 	for _, v := range sp {
	// 		s, _ := strconv.Atoi(v)
	// 		first = append(first, s)
	// 	}
	// 	fmt.Println(first)
	// 	return
	// }
	// fmt.Println(digits)

}
