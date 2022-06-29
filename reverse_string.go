package main

func reverseString(s []byte) []byte {
	leftPtr := 0
	rightPtr := len(s) - 1

	for leftPtr <= rightPtr {
		temp := s[leftPtr]
		s[leftPtr] = s[rightPtr]
		s[rightPtr] = temp

		leftPtr++
		rightPtr--
	}
	return s
}
