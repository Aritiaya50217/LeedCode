package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// res : https://paperbun.org/split-a-string-by-length-in-golang/
func thousandSeparator(n int) string {
	numberStr := strconv.Itoa(n)
	res := ""
	result := []string{}
	if n < 100 {
		return numberStr
	} else if len(numberStr)%3 == 0 {
		sp := 3
		// result := []string{}
		for i := 0; i < len(numberStr); i += sp {
			end := i + sp
			if end > len(numberStr) {
				end = len(numberStr)
			}
			result = append(result, numberStr[i:end])
		}
		res = strings.Join(result, ".")
	} else {
		decimal := math.Floor(float64(n)) / 1000
		res = fmt.Sprintf("%g", decimal)
		return res
	}

	return res
}

func main() {
	n := 2251693
	res := thousandSeparator(n)
	fmt.Println(res)
}
