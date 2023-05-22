package main

import "fmt"

func romanToInt(s string) int {
	numberList := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	num := 0
	totalList := []int{}
	for _, str := range s {
		for key, value := range numberList {
			if string(str) == key {
				totalList = append(totalList, value)
			}
		}
	}
	// for _, t := range totalList {
	// 	num += t
	// }
	for i, _ := range totalList {
		num += totalList[i]
		length := len(totalList) % 2
		if length == 1 {
			if totalList[1] < totalList[2] {
				num = totalList[2] - totalList[1]
				num += totalList[4] - totalList[3]
				num += totalList[6] - totalList[5]
				num += totalList[0]
			}
		} 
	}
	return num
}

func main() {
	input := ""
	fmt.Scan(&input)
	result := romanToInt(input)
	fmt.Println(result)
}
