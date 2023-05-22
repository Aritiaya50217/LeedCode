package main

import (
	"fmt"
	"regexp"
)

func removeDuplicateValues(strList []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strList {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func secondHighest(s string) int {
	re := regexp.MustCompile("[0-9]+")
	strList := re.FindAllString(s, -1)
	list := []string{}
	var lastIndex int
	for _, val := range strList {
		for i := 0; i < len(val); i++ {
			list = append(list, string(val[i]))
		}
	}
	res := removeDuplicateValues(list)
	if len(res) == 1 {
		return -1
	} else if len(res) == 0 {
		return -1
	}
	for i := 0; i < len(res); i++ {
		lastIndex = i
	}

	return lastIndex
}

// TODO
func main() {
	s := "xyz"
	res := secondHighest(s)
	fmt.Println(res)
}
