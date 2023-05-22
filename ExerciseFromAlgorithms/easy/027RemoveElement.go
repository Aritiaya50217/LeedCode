package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		// กรณีที่ค่าเท่ากับ val
		if nums[i] == val {
			// ตัวใดมีค่าเท่ากับ val ให้เพิ่มไป 1
			nums = append(nums[:(i)], nums[(i+1):]...)
			i-- // ลบค่าตัวสุดท้ายไปเรื่อย ๆ
		}
	}
	return len(nums)
	/*	nums = [2,3,3,2]   val = 3
		list[:(i)] : []   list[(i+1):] : [2 3 3 3]
		list[:(i)] : [2]  list[(i+1):] : [3 3 3] => เพื่ม 1 ถ้ามีค่าเท่ากับ val และลบค่าท้ายสุด [2 3 3] => [3 3]
		list[:(i)] : [2]  list[(i+1):] : [3 3] => ไม่มีค่าที่เท่ากับ val ข้ามไปลบตัวท้าย [3]
		list[:(i)] : [2]  list[(i+1):] : [3]  => ไม่มีค่าที่เท่ากับ val ข้ามไปลบตัวท้าย []
		list[:(i)] : [2]  list[(i+1):] : []
		[2 2] => ค่าที่ไม่เท่ากับ val
	*/
}

func main() {
	// https://leetcode.com/problems/remove-element/
	input := bufio.NewScanner(os.Stdin)
	list := []int{}
	for input.Scan() {
		inputStr := input.Text()
		number, _ := strconv.Atoi(inputStr)
		if inputStr != "" {
			list = append(list, number)
		} else {
			break
		}
	}
	last := list[len(list)-1:]
	val := last[0]
	// วิธีที่ 1 ไม่ครอบคลุมทั้งหมดบางค่า error
	// check := []int{}
	// for _, v := range list[:len(list)-1] {
	// 	if v != val {
	// 		check = append(check, v)
	// 	}
	// }
	// fmt.Println(check)

	// วิธีที่ 2
	fmt.Println("--- Use Algorithms ---")
	res := removeElement(list, val)
	fmt.Println(res)
}
