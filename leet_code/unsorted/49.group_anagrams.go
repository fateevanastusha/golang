package main

import (
	"fmt"
	"slices"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}

	seen := map[string]int{}
	for _, str := range strs {
		m := []rune(str)
		slices.Sort(m)
		key := string(m)
		if v, ok := seen[key]; ok {
			res[v] = append(res[v], str)
		} else {
			seen[key] = len(seen)
			res = append(res, []string{str})
		}
	}

	return res
}

func prettyLogGroups(groups [][]string) {
	fmt.Println("Grouped anagrams:")
	for i, g := range groups {
		fmt.Printf(" %d) [%s]\n", i+1, strings.Join(g, ", "))
	}
}

func main() {

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	prettyLogGroups(groupAnagrams(strs))
}
