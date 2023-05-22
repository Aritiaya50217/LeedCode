package main

import "fmt"

func checkZeroOnes(s string) bool {
	countOne, countZero := 0, 0
	for _, v := range s {
		if string(v) == "1" {
			countOne++
		} else {
			countZero++
		}
	}
	if countZero < countOne {
		return true
	}
	return false
}
func main() {
	s := "011000111"
	res := checkZeroOnes(s)
	fmt.Println(res)

}
