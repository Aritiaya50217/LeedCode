package main

import (
	"bufio"
	"fmt"
	"os"
)

func findRotation(mat [][]int, target [][]int) bool {
	count1, count2 := 0, 0
	for i := 0; i < len(mat); i++ {
		for _, v := range mat[i] {
			if v == 1 {
				count1++
			}
		}
		for _, v := range target[i] {
			if v == 1 {
				count2++
			}
		}
	}
	if count1 == count2 {
		return true
	}
	return false
}

func findRotation2(mat [][]int, target [][]int) bool {
	m1, m2, t1, t2 := 0, 0, 0, 0
	for _, v := range mat {
		m1 = v[0]
		m2 = v[1]
		for _, val := range target {
			t1 = val[0]
			t2 = val[1]
		}
	}
	output := t1 == m2 && t2 == m1
	if output == true {
		return true
	}
	return false

}

func areIntSlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// TODO 

func main() {
	// num1 := []int{0, 1}
	// num2 := []int{1, 1}
	// mat := make([][]int, 0)
	// target := make([][]int, 0)
	// mat = append(mat, num1, num2)
	// target = append(target, num2, num1)
	// fmt.Printf("mat : %v , target : %v\n", mat, target)

	input := bufio.NewScanner(os.Stdin)
	list := []string{}
	// num1 := []int{}
	// num2 := []int{}
	for input.Scan() {
		inputStr := input.Text()
		if inputStr != "" {
			list = append(list, inputStr)
		} else {
			break
		}
	}
	// ex 9 => [1,2,3] [4,5,6] [7,8,9]
	length := len(list)
	ranges := 0
	for length/2 == 0 {
		ranges++
	}
	fmt.Println(ranges)

	// r := length / 2
	// ranges := length / r
	// fmt.Println(length, r, ranges)
	// for i := 0; i < len(list); i++ {
	// 	number, _ := strconv.Atoi(list[i])
	// 	if len(num1) < ranges {
	// 		num1 = append(num1, number)
	// 	} else if len(num1) == ranges {
	// 		num2 = append(num2, number)
	// 	}
	// }
	// fmt.Println(num1, num2)

}
