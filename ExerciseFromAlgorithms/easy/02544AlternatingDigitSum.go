package main

import (
	"fmt"
	"math"
	"strconv"
)

func alternateDigitSum(n int) int {
	str := strconv.Itoa(n)
	num := 0.0
	list := []float64{}
	for i := 0; i < len(str); i++ {
		num, _ = strconv.ParseFloat(string(str[i]), 64)
		list = append(list, num)
	}
	number := 0.0
	s := []string{}
	for i, v := range list {
		if i%2 != 0 {
			// ref :https://stackoverflow.com/questions/13804255/negative-zero-literal-in-golang
			number = math.Copysign(float64(v), -1)
			ss := fmt.Sprint(number)
			s = append(s, ss)
		} else {
			ss := fmt.Sprint(v)
			s = append(s, ss)
		}
	}
	sum := 0
	for _, val := range s {
		n, _ := strconv.Atoi(val)
		sum += n
	}
	return sum
}
func main() {
	n := 886996
	res := alternateDigitSum(n)
	fmt.Println(res)

}
