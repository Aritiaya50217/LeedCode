package main

import "fmt"

func sortEvenOdd(nums []int) []int {
	oddList, evenList := []int{}, []int{}
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			evenList = append(evenList, i)
		} else {
			oddList = append(oddList, i)
		}
	}

	// ref : https://linuxhint.com/golang-reverse-slice/
	revEven, revOdd, list, res := []int{}, []int{}, []int{}, []int{}
	for i, _ := range evenList {
		revEven = append(revEven, evenList[len(evenList)-1-i])
		revOdd = append(revOdd, oddList[len(oddList)-1-i])
	}
	for j := 0; j < len(revEven); j++ {
		list = append(list, revEven[j], revOdd[j])
	}

	for k := 0; k < len(nums); k++ {
		res = append(res, nums[list[k]])
	}
	return res
}

// TODO
func main() {
	nums := []int{36, 45, 32, 31, 15, 41, 9, 46, 36, 6, 15, 16, 33, 26, 27, 31, 44, 34}
	// index := []int{6, 10, 4, 1, 9, 5, 14, 17, 2, 15, 12, 15, 0, 13, 8, 11, 16, 9}
	oddList, evenList := []int{}, []int{}

	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			evenList = append(evenList, i)
		} else {
			oddList = append(oddList, i)
		}
	}
	fmt.Println(evenList, oddList)

	// revEven, revOdd, list, res := []int{}, []int{}, []int{}, []int{}
	// for i, _ := range evenList {
	// 	revEven = append(revEven, evenList[len(evenList)-1-i])
	// 	revOdd = append(revOdd, oddList[len(oddList)-1-i])
	// }
	// for j := 0; j < len(revEven); j++ {
	// 	list = append(list, revEven[j], revOdd[j])
	// }

	// for k := 0; k < len(nums); k++ {
	// 	res = append(res, nums[list[k]])
	// }
	// fmt.Println(res)

	// nums = [36,45,32,31,15,41,9,46,36,6,15,16,33,26,27,31,44,34]
	// เฉลย = [9,46,15,45,15,41,27,34,32,31,33,31,36,26,36,16,44,6]
	// index : [6,10,4,1,9,5,14,17,2,15,12,15,0,13,8,11,16,9]

}
