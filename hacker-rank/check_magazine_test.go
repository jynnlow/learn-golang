package main

import "testing"

type checkMagazineStruct struct {
	magazine []string
	note     []string
	answer   bool
}

func Test_checkMagazine(t *testing.T) {
	testCheckMagazineArray := []checkMagazineStruct{
		{
			magazine: []string{"give", "Me", "one", "grand", "today", "night"},
			note:     []string{"give", "one", "today", "me", "grand", "night"},
			answer:   true,
		},
		{
			magazine: []string{"o", "l", "x", "imjaw", "bee", "khmla", "v", "o", "v", "o", "imjaw", "l", "khmla", "imjaw", "x"},
			note:     []string{"imjaw", "l", "khmla", "x", "imjaw", "o", "l", "l", "o", "khmla", "v", "bee", "o", "o", "imjaw", "imjaw", "o"},
			answer:   false,
		},
	}

	for _, testCheckMagazine := range testCheckMagazineArray {
		if checkMagazine(testCheckMagazine.magazine, testCheckMagazine.note) != testCheckMagazine.answer {
			t.Error("Wrong Answer")
		}
	}
}
