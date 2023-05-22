package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

// ref : https://www.tutorialspoint.com/write-a-golang-program-to-find-duplicate-elements-in-a-given-array
func duplicateInArray(arr []int) int {
	visited := make(map[int]bool, 0)
	for i := 0; i < len(arr); i++ {
		if visited[arr[i]] == true {
			return arr[i]
		} else {
			visited[arr[i]] = true
		}
	}
	return -1
}

func areNumbersAscending(s string) bool {
	re := regexp.MustCompile("[0-9]+")
	strList := re.FindAllString(s, -1)
	array := []int{}
	for _, v := range strList {
		n, _ := strconv.Atoi(v)
		array = append(array, n)
	}

	check := duplicateInArray(array)
	if check == -1 {
		list := []int{}
		for _, v := range strList {
			number, _ := strconv.Atoi(v)
			list = append(list, number)
		}
		// ref : https://surajsharma.net/blog/golang-ints-are-sorted-func
		return sort.IntsAreSorted(list)
	} else {
		return false
	}
}

func main() {
	s := "hello world 5 x 5"
	res := areNumbersAscending(s)
	fmt.Println(res)
}
