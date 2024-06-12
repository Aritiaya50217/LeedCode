package main

import "fmt"

func sum(num1 int, num2 int) int {
	num1 += num2
	return num1
}
func main() {
	num1, num2 := -10, 4
	res := sum(num1, num2)
	fmt.Println(res)
}
