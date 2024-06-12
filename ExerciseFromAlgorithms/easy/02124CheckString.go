package main

import (
	"fmt"
	"sort"
	"strings"
)

// ref :https://www.geeksforgeeks.org/how-to-sort-a-slice-of-strings-in-golang/
func checkString(s string) bool {
	// เรียงตาม digit
	// = true , < false
	sp := strings.Split(s, "")
	check := sort.StringsAreSorted(sp)
	return check
}

func main() {
	s := "abab"
	res := checkString(s)
	fmt.Println(res)

	// a := 1
	// b := 2

	// Texta := "a"
	// // TextA := "A"

	// Textb := "b"
	// // TextB := "B"

	// inPut := "aaabaab"
	// inPut[1] = a
	// sss := 1

	// if sss >= check
}
