package main

import "fmt"

func maxValue(arrays []int, length int) int {
	max := 0
	for i := 0; i < length-1; i++ {
		if arrays[i] > arrays[i+1] { // swap
			var temp = arrays[i]
			arrays[i] = arrays[i+1]
			arrays[i+1] = temp
		}
	}
	max = arrays[length-1]
	return max
}

func minValue(arrays []int, length int) int {
	min := 0
	for i := 0; i < length-1; i++ {
		if arrays[i] < arrays[i+1] { // swap
			var temp = arrays[i]
			arrays[i] = arrays[i+1]
			arrays[i+1] = temp
		}
	}
	min = arrays[length-1]
	return min
}

func average(salary []int) float64 {
	max := maxValue(salary, len(salary))
	min := minValue(salary, len(salary))
	list := []int{}
	total := 0.0
	for _, v := range salary {
		if v != max && v != min {
			list = append(list, v)
		}
	}
	for _, v := range list {
		total += float64(v)
	}
	return float64(total) / float64(len(list))
}

func main() {
	salary := []int{8000, 9000, 2000, 3000, 6000, 1000}
	res := average(salary)
	fmt.Printf("use func : %f\n", res)

}
