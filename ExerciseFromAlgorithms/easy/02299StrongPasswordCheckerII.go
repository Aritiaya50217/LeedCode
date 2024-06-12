package main

import (
	"fmt"
	"regexp"
)

func strongPasswordCheckerII(password string) bool {
	secure := true
	tests := []string{".{7,}", "[a-z]", "[A-Z]", "[0-9]", "[^\\d\\w]"}
	for _, test := range tests {
		t, _ := regexp.MatchString(test, password)
		if !t {
			secure = false
			break
		}
	}
	return secure
}

// TODO
func main() {
	// password := "11A!A!Aa"
	// secure := true
	reg := []string{".{7,}", "[a-z]", "[A-Z]", "[0-9]", "[^\\d\\w]", "[!@#$%^&*()-+]"}
	for _, v := range reg {
		fmt.Println(v)
	}

	// for i, val := range reg {
	// 	t, _ := regexp.MatchString(val, password)
	// 	if !t {
	// 		secure = false
	// 		break
	// 	}
	// }
	// fmt.Println(secure)

}
