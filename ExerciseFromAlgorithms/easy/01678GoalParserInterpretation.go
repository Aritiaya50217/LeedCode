package main

import (
	"fmt"
	"strings"
)

// ref : https://blog.boot.dev/golang/replace-strings-golang/
func interpret(command string) string {
	replacer := strings.NewReplacer("()", "o", "(", "", ")", "")
	res := replacer.Replace(command)
	return res
}

func main() {
	input := ""
	fmt.Scan(&input)
	res := interpret(input)
	fmt.Println(res)
}
