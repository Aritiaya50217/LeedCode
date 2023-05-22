package main

import (
	"fmt"
	"strings"
)
// TODO 
func modifyString(s string) string {
	replacer := strings.NewReplacer("??", "ac", "?", "a")
	res := replacer.Replace(s)
	return res
}
func main() {
	s := "j?qg??b"
	res := modifyString(s)
	fmt.Println(res)

}
