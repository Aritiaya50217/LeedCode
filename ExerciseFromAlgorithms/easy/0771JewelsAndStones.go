package main

import "fmt"

// หาตัวที่ซ้ำกัน
func numJewelsInStones(jewels string, stones string) int {
	runeJ := []rune(jewels)
	runeS := []rune(stones)
	count := 0
	for i := 0; i < len(runeJ); i++ {
		for j := 0; j < len(runeS); j++ {
			// check
			if runeJ[i] == runeS[j] {
				count++
			}
		}
	}
	return count
}

func main() {
	jewels := ""
	stones := ""
	fmt.Scan(&jewels, &stones)
	res := numJewelsInStones(jewels, stones)
	fmt.Println(res)

}
