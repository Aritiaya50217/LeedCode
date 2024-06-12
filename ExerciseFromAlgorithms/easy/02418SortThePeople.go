package main

import (
	"fmt"
	"sort"
)

type Data struct {
	Names   string
	Heights int
}

func sortPeople(names []string, heights []int) []string {
	var people *Data
	h := []int{}
	for i := 0; i < len(names); i++ {
		people = &Data{
			Names:   names[i],
			Heights: heights[i],
		}
		h = append(h, people.Heights)
		sort.Slice(h, func(i, j int) bool {
			return h[i] > h[j]
		})
	}
	namesList := []string{}
	for j := 0; j < len(heights); j++ {
		for k := 0; k < len(h); k++ {
			if heights[j] == h[k] {
				namesList = append(namesList, names[k])
			}
		}
	}
	return namesList
}

// TODO
func main() {
	names := []string{"Mary", "John", "Emma"}
	heights := []int{180, 165, 170}

	// var people *Data
	// h := []int{}
	// for i := 0; i < len(names); i++ {
	// 	people = &Data{
	// 		Names:   names[i],
	// 		Heights: heights[i],
	// 	}
	// 	h = append(h, people.Heights)
	// 	sort.Slice(h, func(i, j int) bool {
	// 		return h[i] > h[j]
	// 	})
	// }
	// namesList := []string{}
	// for j := 0; j < len(heights); j++ {
	// 	for k := 0; k < len(h); k++ {
	// 		if heights[j] == h[k] {
	// 			namesList = append(namesList, names[k])
	// 		}
	// 	}
	// }
	// fmt.Println(namesList)

	res := sortPeople(names, heights)
	fmt.Println(res)

}
