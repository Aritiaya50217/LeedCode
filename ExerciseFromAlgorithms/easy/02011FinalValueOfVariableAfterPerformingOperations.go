package main

import (
	"fmt"
	"strconv"
	"strings"
)

func interpret(command string) string {
	replacer := strings.NewReplacer("--X", "-1", "X--", "-1", "++X", "1", "X++", "1")
	res := replacer.Replace(command)
	return res
}
func finalValueAfterOperations(operations []string) int {
	sum := 0
	for _, v := range operations {
		res := interpret(v)
		num, _ := strconv.Atoi(res)
		sum += num
	}
	return sum
}
func main() {
	operations := []string{"--X", "++X", "X++"}
	res := finalValueAfterOperations(operations)
	fmt.Println(res)

}
