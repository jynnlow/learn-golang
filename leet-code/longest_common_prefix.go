package main

import "sort"

func longestCommonPrefix(strs []string) string {
	var longestPrefix string = ""

	if len(strs) > 0 {
		//["flight", "flow", "flower"]
		sort.Strings(strs)
		//first = "flight"
		first := string(strs[0])
		//last = "flower"
		last := string(strs[len(strs)-1])

		for i := 0; i < len(first); i++ {
			if string(last[i]) == string(first[i]) {
				longestPrefix += string(last[i])
			} else {
				break
			}
		}
	}
	return longestPrefix
}
