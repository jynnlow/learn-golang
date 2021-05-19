package main

import (
	"strings"
)

func twoStrings(s1 string, s2 string) string {
	splitedFirstString := strings.Split(s1,"")
	splitedSecondString := strings.Split(s2, "")
	sharedCommonString := false

	stringMap := make(map[string]int)
	for i:= 0; i<len(s1); i++{
		stringMap[splitedFirstString[i]] = 1
	}

	for i:= 0; i<len(s2); i++{
		if _, ok := stringMap[splitedSecondString[i]]; ok{
			sharedCommonString = true
			break
		}else {
			sharedCommonString = false
		}
	}

	if sharedCommonString {
		return "YES"
	}else{
		return "NO"
	}
}
