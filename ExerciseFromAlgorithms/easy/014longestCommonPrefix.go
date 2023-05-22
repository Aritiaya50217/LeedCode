package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func longestCommonPrefix(strs []string) string {
	var longestPrefix string = ""
	var endPrefix = false
	if len(strs) > 0 {
		sort.Strings(strs)
		first := string(strs[0])
		last := string(strs[len(strs)-1])

		for i := 0; i < len(first); i++ {
			if !endPrefix && string(last[i]) == string(first[i]) {
				longestPrefix += string(last[i])
			} else {
				endPrefix = true
			}
		}
	}
	return longestPrefix
}
func main() {
	input := bufio.NewScanner(os.Stdin)
	list := []string{}
	for input.Scan() {
		inputStr := input.Text()
		list = append(list, inputStr)
		if len(list) == 3 {
			break
		}
	}
	res := longestCommonPrefix(list)
	fmt.Println(res)
}
