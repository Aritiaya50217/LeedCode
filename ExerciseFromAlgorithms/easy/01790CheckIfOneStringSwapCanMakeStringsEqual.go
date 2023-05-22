package main

import "fmt"

// func areAlmostEqual(s1 string, s2 string) bool {
// 	count, count2 := 0, 0
// 	for _, v1 := range s1 {
// 		count += int(v1)
// 	}
// 	for _, v2 := range s2 {
// 		count2 += int(v2)
// 	}
// 	if count == count2 {
// 		return true
// 	}

// 	return false
// }

func areAlmostEqual(s1 string, s2 string) bool {
	first, second, count := -1, -1, 0

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			count++
			if count > 2 {
				return false
			}

			if first == -1 && second == -1 {
				first = i
			} else {
				second = i
			}
		}
	}

	if count == 0 {
		return true
	}

	if first != -1 && second != -1 && s1[first] == s2[second] && s1[second] == s2[first] {
		return true
	}

	return false

}

func main() {
	s1, s2 := "abcd", "dcba"
	res := areAlmostEqual(s1, s2)
	fmt.Println(res)

}
