package main

import (
	"fmt"
	"strconv"
)

//func main(){
//	fmt.Println(quicksort([]int{6,4,1}))
//}

func countSwaps(a []int32) {
	counter := 0
	for i := 0; i < len(a) - 1; i ++{
		for j := i + 1; j < len(a); j++{
			if a[i] > a[j]{
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
				counter += 1
			}
		}
	}
	fmt.Println("Array is sorted in " + strconv.Itoa(counter) + " swaps.")
	fmt.Println("First Element: " + strconv.Itoa(int(a[0])))
	fmt.Println("Second Element: " + strconv.Itoa(int(a[len(a)-1])))
}
