package main

import (
	"fmt"
	"strconv"
)

// ref :https://stackoverflow.com/questions/25686109/split-string-by-length-in-golang
// splitBy splits a string s by int n
func splitBy(s string, n int) []string {
	var ss []string
	for i := 1; i < len(s); i++ {
		if i%n == 0 {
			ss = append(ss, s[:i])
			s = s[i:]
			i = 1
		}
	}
	ss = append(ss, s)
	return ss
}

func getLucky(s string, k int) int {
	alphabet := map[int]string{1: "a", 2: "b", 3: "c", 4: "d",
		5: "e", 6: "f", 7: "g", 8: "h", 9: "i", 10: "j", 11: "k",
		12: "l", 13: "m", 14: "n", 15: "o", 16: "p", 17: "q", 18: "r",
		19: "s", 20: "t", 21: "u", 22: "v", 23: "w", 24: "x", 25: "y",
		26: "z"}
	count := 0
	switch k {
	case 1:
		for i, v := range alphabet {
			for _, val := range s {
				if string(val) == v {
					count += i
				}
			}
		}
		return count

	default:
		count := 0
		switch k {
		case 1:
			for i, v := range alphabet {
				for _, val := range s {
					if string(val) == v {
						count += i
					}
				}
			}
			fmt.Println(count)
		default:
			char := []string{}
			list := []string{}
			for i, v := range alphabet {
				for _, val := range s {
					if string(val) == v {
						indexStr := strconv.Itoa(i)
						char = splitBy(indexStr, 1)
						for _, s := range char {
							list = append(list, s)
						}
					}
				}
			}
			for _, v := range list {
				nums, _ := strconv.Atoi(string(v))
				count += nums
			}
			return count

		}

	}
	return count
}

// TODO
func main() {
	alphabet := map[int]string{1: "a", 2: "b", 3: "c", 4: "d",
		5: "e", 6: "f", 7: "g", 8: "h", 9: "i", 10: "j", 11: "k",
		12: "l", 13: "m", 14: "n", 15: "o", 16: "p", 17: "q", 18: "r",
		19: "s", 20: "t", 21: "u", 22: "v", 23: "w", 24: "x", 25: "y",
		26: "z"}

	s := "qhquvppzooyt"
	k := 6

	count := 0
	switch k {
	case 1:
		for i, v := range alphabet {
			for _, val := range s {
				if string(val) == v {
					count += i
				}
			}
		}
		fmt.Println(count)
	default:
		char := []string{}
		list := []string{}
		for i, v := range alphabet {
			for _, val := range s {
				if string(val) == v {
					indexStr := strconv.Itoa(i)
					char = splitBy(indexStr, 1)
					for _, s := range char {
						list = append(list, s)
					}
				}
			}
		}
		for _, v := range list {
			nums, _ := strconv.Atoi(string(v))
			count += nums
		}
		fmt.Println(count)

	}

}
