package main

import "fmt"

func main() {
	fmt.Println(makeAnagram("fcrxzwscanmligyxyvym", "jxwtrhvujlmrpdoqbisbwhmgpmeoke"))
}

func makeAnagram(a string, b string) int32 {
	mapA := make (map[string]int)
	var totalMapCount = int32(len(a))
	var deletion int32

	for i := 0; i < len(a); i++{
		mapA [string(a[i])] += 1
	}

	for i := 0; i < len(b); i++{
		if _, ok := mapA[string(b[i])]; ok{
			mapA[string(b[i])] -= 1
			totalMapCount -= 1
			if mapA[string(b[i])] == 0{
				delete(mapA, string(b[i]))
			}
		}else {
			deletion += 1
		}
	}
	deletion += totalMapCount
	return  deletion
}
