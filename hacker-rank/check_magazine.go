package main

import "fmt"

func checkMagazine(magazine []string, note []string) bool {
	magazineMap := make(map[string]int)
	canFormed := true

	//this for loop is to check if there is any element repeated in the magazine string array more than one time in the map
	for i := 0; i < len(magazine); i++ {
		if _, ok := magazineMap[magazine[i]]; ok {
			magazineMap[magazine[i]] += 1
		} else {
			magazineMap[magazine[i]] = 1
		}
	}
	fmt.Println(magazineMap)

	//this for loop is to check the note string array matched with the magazine map
	for i := 0; i < len(note); i++ {
		if _, ok := magazineMap[note[i]]; ok {
			magazineMap[note[i]] -= 1
			if magazineMap[note[i]] == 0 {
				delete(magazineMap, note[i])
			}
		} else {
			canFormed = false
		}
	}
	fmt.Println(magazineMap)

	return canFormed
}
