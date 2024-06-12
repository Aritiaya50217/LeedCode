package main

import (
	"fmt"
)

func convertTemperature(celsius float64) []float64 {
	kelvin := celsius + 273.15
	fahrenheit := celsius*1.80 + 32.00
	list := []float64{kelvin, fahrenheit}
	return list
}

// TODO
func main() {
	celsius := 122.11
	res := convertTemperature(celsius)
	fmt.Println(res)
}
