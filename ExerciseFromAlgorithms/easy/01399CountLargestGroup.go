package main

import "fmt"

func countLargestGroup(n int) int {
	maxGroup, maxCount := 0, 0
	count := make(map[int]int)
	for i := 1; i <= n; i++ {
		num := i
		sum := 0
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		count[sum]++
		if count[sum] == maxCount {
			maxGroup++
		} else if count[sum] > maxCount {
			maxGroup = 1
			maxCount = count[sum]
		}
	}
	return maxGroup
}

func countLargestGroup2(n int) int {
	list := []int{}
	list2 := []int{}
	list3 := []int{}
	length, total := 0, 0
	for i := 1; i <= n; i++ {
		total = i / 10
		if total == 0 {
			list = append(list, total)
			length = len(list)
		} else {
			list2 = append(list2, total)
			length = len(list2)
		}
	}
	for _, v := range list2 {
		r := max(list2, len(list2))
		if v == r {
			list3 = append(list3, v)
			length = len(list3)
		}
	}
	return length
}

func max(array []int, length int) int {
	for i := 0; i < length-1; i++ {
		if array[i] > array[i+1] { // swap
			var temp = array[i]
			array[i] = array[i+1]
			array[i+1] = temp
		}
	}
	var maxValue = array[length-1]
	return maxValue
}

func main() {
	input := 0
	fmt.Scan(&input)
	res := countLargestGroup(input)
	fmt.Println("use func :", res)

	// res2 := countLargestGroup2(input)
	// fmt.Println(res2)

	list := []int{}
	list2 := []int{}
	list3 := []int{}
	length, total := 0, 0
	for i := 1; i <= input; i++ {
		total = i / 10
		if total == 0 {
			list = append(list, total)
			length = len(list)
		} else {
			list2 = append(list2, total)
			length = len(list2)
		}
		fmt.Printf("list : %v\nlist2 :%v\n", list, list2)
	}
	for _, v := range list2 {
		r := max(list2, len(list2))
		if v == r {
			list3 = append(list3, v)
			length = len(list3)
		}
	}
	fmt.Println(length)

}
