package main

func intersectionOfTwoArrays(nums1 []int, nums2 []int) []int {
	intersection := make([]int, 0)
	comparisonMap := make(map[int]int)

	for _, num1 := range nums1 {
		_, exist := comparisonMap[num1]
		if exist {
			comparisonMap[num1] += 1
		} else {
			comparisonMap[num1] = 1
		}
	}

	for _, num2 := range nums2 {
		_, exist := comparisonMap[num2]
		if exist {
			comparisonMap[num2] -= 1
			intersection = append(intersection, num2)
			if comparisonMap[num2] == 0 {
				delete(comparisonMap, num2)
			}
		}
	}
	return intersection
}
