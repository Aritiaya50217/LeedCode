package main

import "fmt"

func isIsomorphic(s, t string) bool {
	m1 := make([]int, 256)
	m2 := make([]int, 256)
	for i := 0; i < len(s); i++ {
		if m1[int(s[i])] != m2[int(t[i])] {
			return false
		}
		i2 := i + 1
		m1[int(s[i])] = i2
		m2[int(t[i])] = i2
	}

	return true
}

func main() {
	s1 := ""
	s2 := ""
	fmt.Scan(&s1, &s2)
	res := isIsomorphic(s1, s2)
	fmt.Println(res)
}
