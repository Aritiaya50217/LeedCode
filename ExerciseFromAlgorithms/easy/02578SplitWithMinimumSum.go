package main

import (
	"fmt"
	"sort"
	"strconv"
)

func splitNum(num int) int {
	str := strconv.Itoa(num)
	evenList, oddList := []int{}, []int{}
	for i := 0; i < len(str); i++ {
		number, _ := strconv.Atoi(string(str[i]))
		if i%2 == 0 {
			evenList = append(evenList, number)
		} else {
			oddList = append(oddList, number)
		}
		// sort
		sort.Ints(oddList)
		sort.Ints(evenList)
	}
	oddStr, evenStr := "", ""
	list := []string{}
	for j := 0; j < len(oddList); j++ {
		os := strconv.Itoa(oddList[j])
		oddStr += os
		es := strconv.Itoa(evenList[j])
		evenStr += es
	}
	list = append(list, evenStr, oddStr)
	total := 0
	for _, v := range list {
		s, _ := strconv.Atoi(v)
		total += s
	}

	return total
}
func main() {
	num := 687
	str := strconv.Itoa(num)
	evenList, oddList := []int{}, []int{}
	for i := 0; i < len(str); i++ {
		number, _ := strconv.Atoi(string(str[i]))
		if i%2 == 0 {
			evenList = append(evenList, number)
		} else {
			oddList = append(oddList, number)
		}
		// sort
		sort.Ints(oddList)
		sort.Ints(evenList)
	}
	fmt.Println(oddList, evenList)
	oddStr, evenStr := "", ""
	list := []string{}
	for j := 0; j < len(oddList); j++ {
		os := strconv.Itoa(oddList[j])
		oddStr += os
		es := strconv.Itoa(evenList[j])
		evenStr += es
	}
	list = append(list, evenStr, oddStr)
	total := 0
	for _, v := range list {
		s, _ := strconv.Atoi(v)
		total += s
	}

	fmt.Println(total)

	// res := splitNum(num)
	// fmt.Println(res)
}
