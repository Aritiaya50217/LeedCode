package main

import (
	"fmt"
	"strings"
)

func defangIPaddr(address string) string {
	res := strings.Replace(address, ".", "[.]", -1)
	return res
}

func main() {
	input := ""
	fmt.Scan(&input)
	res := defangIPaddr(input)
	fmt.Println(res)
}
