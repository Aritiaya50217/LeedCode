package main

import (
	"fmt"
)

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	// set up stack and map
	st := []rune{}
	// map[key][value]
	open := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, r := range s {
		// ถ้า open[r] = value ใน open[] ให้เพิ่มลงใน st[]
		// ex open["("] => value ")"
		if key, value := open[r]; value {
			st = append(st, key)
			continue
		}
		// check to make sure the stack isn't empty
		length := len(st) - 1
		if length < 0 || r != st[length] {
			return false
		}

		// take the last element off the stack
		st = st[:length]
	}
	return len(st) == 0
}

func main() {
	input := ""
	fmt.Scan(&input)
	// res := isValid(input)
	// fmt.Println(res)

	if len(input)%2 != 0 {
		fmt.Println(false)
	}
	st := []rune{}
	open := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, r := range input {
		fmt.Println("r :", string(r))
		if key, value := open[r]; value {
			// fmt.Println("open[r] :", string(open[r])) // value
			st = append(st, key)
			continue
		}
		// check to make sure the stack isn't empty
		length := len(st) - 1
		if length < 0 || r != st[length] {
			fmt.Println(false)
			return
		}
		// take the last element off the stack
		st = st[:length]
	}
	fmt.Println(len(st) == 0)

}
