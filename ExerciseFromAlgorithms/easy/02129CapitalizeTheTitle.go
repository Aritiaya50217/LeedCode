package main

import (
	"fmt"
	"strings"
)

func capitalizeTitle(title string) string {
	title = strings.ToLower(title)
	list := strings.Split(title, " ")
	res := ""
	strList := []string{}
	for _, v := range list {
		for i := 0; i < len(v); i++ {
			if len(v) <= 2 {
				str := string(v)
				strList = append(strList, str)
				break
			} else {
				if i == 0 {
					upper := strings.ToUpper(string(v[i]))
					str := strings.Replace(string(v), string(v[i]), upper, 1)
					strList = append(strList, str)
				}
			}
		}
	}
	res = strings.Join(strList, " ")
	return res[0:]
}

func main() {
	title := "First leTTeR of EACH Word"
	res := capitalizeTitle(title)
	fmt.Println(res)

}
