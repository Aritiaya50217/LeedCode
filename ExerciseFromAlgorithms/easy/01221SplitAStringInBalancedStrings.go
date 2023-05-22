package main

import (
	"fmt"
	"strings"
)

func balancedStringSplit(s string) int {
	return 0
}
func main() {
	input := ""
	fmt.Scan(&input)
	str := "RL"
	for _, v := range input {
		if strings.Contains(str, string(v)) {
			fmt.Println("Yes")
		} else {
			continue
		}
	}

}
