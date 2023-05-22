package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func sortSentence(s string) string {
	return ""
}

// TODO
func main() {
	s := "is2 sentence4 This1 a3"
	re := regexp.MustCompile("[0-9]+")
	indexAll := re.FindAllString(s, -1)
	indexList := []int{}
	for i := 0; i < len(indexAll); i++ {
		n, _ := strconv.Atoi(indexAll[i])
		indexList = append(indexList, n)
	}

	word := regexp.MustCompile(`[^a-zA-Z]+`).ReplaceAllString(s, " ")
	sp := strings.Split(word[:len(word)-1], " ")
	// fmt.Println(sp, indexList)

	res := make([]string, len(sp))
	for i := 0; i < len(sp); i++ {
		res[indexList[i]] = sp[i]
	}
	fmt.Println(res)

}
