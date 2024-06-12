package main

import "fmt"

func isCircularSentence(sentence string) bool {
	if sentence[0] != sentence[len(sentence)-1] {
		return false
	}

	for i := 2; i < len(sentence); i++ {
		if sentence[i-1] == ' ' && sentence[i-2] != sentence[i] {
			return false
		}
	}

	return true
}
func main() {
	sentence := "MuFoevIXCZzrpXeRmTssj lYSW U jM"
	res := isCircularSentence(sentence)
	fmt.Println(res)

}
