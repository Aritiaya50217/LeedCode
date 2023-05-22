package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ref : https://stackoverflow.com/questions/13737745/split-a-string-on-whitespace-in-go
func lengthOfLastWord(s string) int {
	sp := strings.Fields(s)
	length := 0
	for _, v := range sp[len(sp)-1:] {
		length = len(v)
	}
	return length
}
func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		inputStr := input.Text()
		res := lengthOfLastWord(inputStr)
		fmt.Println(res)
	}

}
