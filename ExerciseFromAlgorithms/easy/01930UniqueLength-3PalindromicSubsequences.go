package main

import (
	"fmt"
	"math"
)

// ref : https://www.tutorialspoint.com/golang-program-to-find-the-duplicate-characters-in-the-string
func countPalindromicSubsequence(s string) int {
	counts := make(map[rune]int)
	for _, v := range s {
		counts[v]++
	}
	for _, count := range counts {
		if count > 1 {
			return count
		}
	}
	return 0
}

// วิธีที่ 2
func countPalindromicSubsequence2(s string) int {
	counter := 0
	first := make([]int, 26)
	last := make([]int, 26)
	for i := range first {
		first[i] = math.MaxInt32
	}

	for i := range last {
		last[i] = -1
	}

	for i, v := range s {
		if first[v-'a'] == math.MaxInt32 {
			first[v-'a'] = i
		} else {
			last[v-'a'] = i
		}
	}
	for i := 0; i < 26; i++ {
		magic := make(map[byte]struct{}, 26)
		for index := first[i] + 1; index < last[i]; index++ {
			magic[s[index]] = struct{}{}
		}
		counter += len(magic)
	}
	return counter
}

func main() {
	input := ""
	fmt.Scan(&input)
	res := countPalindromicSubsequence2(input)
	fmt.Println(res)
	
	counter := 0
	first := make([]int, 26)
	last := make([]int, 26)
	for i := range first {
		first[i] = math.MaxInt32
	}

	for i := range last {
		last[i] = -1
	}

	for i, v := range input {
		if first[v-'a'] == math.MaxInt32 {
			first[v-'a'] = i
		} else {
			last[v-'a'] = i
		}
	}
	for i := 0; i < 26; i++ {
		magic := make(map[byte]struct{}, 26)
		for index := first[i] + 1; index < last[i]; index++ {
			magic[input[index]] = struct{}{}
		}
		counter += len(magic)
	}
	fmt.Println(counter)

}
