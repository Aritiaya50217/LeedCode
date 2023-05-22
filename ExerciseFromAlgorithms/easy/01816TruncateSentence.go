package main

import (
	"fmt"
	"strings"
)

func truncateSentence(s string, k int) string {
	sp := strings.Split(s, " ")
	res := ""
	for _, v := range sp[:k] {
		res += v + " "
	}
	res = res[:len(res)-1]
	return res
}

func main() {
	s := "Hello how are you Contestant"
	k := 4
	res := truncateSentence(s,k)
	fmt.Println(res)
	
}
