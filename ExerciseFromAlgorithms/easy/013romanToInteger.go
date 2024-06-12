package main

import (
	"strings"
)

//TODO
func romanToInt(s string) int {
	roman := map[string]int{"I": 1, "V": 5, "X": 10,
		"L": 50, "C": 100, "D": 500, "M": 1000}
	sp := strings.Split(s, "")
	count := 0
	for _, v := range sp {
		for key, value := range roman {
			if v == key {
				count += value
			}
		}
	}
	return count
}
func main() {
	s := "MCMXCIV"
	roman := map[string]int{"I": 1, "V": 5, "X": 10,
		"L": 50, "C": 100, "D": 500, "M": 1000}
	sp := strings.Split(s, "")
	list := []int{}
	for i := 0; i < len(sp); i++ {
		for key, value := range roman {
			if key == sp[i] {
				list = append(list, value)
			}
		}
	}

	// count := 0
	// for _, v := range sp {
	// 	for key, value := range roman {
	// 		if v == key {
	// 			count += value
	// 		}
	// 	}
	// }
	// fmt.Println(count)

}
