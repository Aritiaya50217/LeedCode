package main

import (
	"fmt"
	"strconv"
)

func percentageLetter(s string, letter byte) int {
	count := 0.0
	for _, v := range s {
		if string(v) == string(letter) {
			count++
		} else {
			return 0
		}
	}

	countFloat := fmt.Sprintf("%0.2f", count)
	lengthFloat := fmt.Sprintf("%0.2f", float64(len(s)))

	num1, _ := strconv.ParseFloat(countFloat, 64)
	num2, _ := strconv.ParseFloat(lengthFloat, 64)

	total := (num1 / num2) * 100
	res := int(total)
	return res
}
func main() {
	s, letter := "jjjj", "k"
	list := []byte(letter)
	res := percentageLetter(s, list[0])
	fmt.Println(res)
}
