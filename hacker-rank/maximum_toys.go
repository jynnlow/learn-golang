package main

import (
	"fmt"
	"math/rand"
)

func main(){
	fmt.Println(maximumToys([]int32{1, 12, 5, 111, 200, 1000, 10},50))
}

func maximumToys(prices []int32, k int32) int32 {
	var itemCount int
	var subtotal int32

	prices = quicksort(prices)

	for i := 0; i < len(prices); i++{
		subtotal += prices[i]
		if subtotal > k {
			itemCount = i
			break
		}
	}

	return int32(itemCount)
}

func quicksort(a []int32) []int32 {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}

