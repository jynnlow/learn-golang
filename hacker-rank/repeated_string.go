package main

/*
This is Kahwei's logic and optimised solution
- findA and findAWithinSubstring functions
 */
func findA(substring string, maxlength int) int {
	substringMap := findAWithinSubstring(substring)
	if maxlength % len(substring) == 0 {
		return substringMap[len(substring)-1] * (maxlength/len(substring))
	}else {
		remainder := maxlength % len(substring)
		countOfA := substringMap[len(substring)-1] * (maxlength/len(substring))
		return countOfA + substringMap[remainder-1]
	}
}

func findAWithinSubstring(substring string) map[int]int {
	aMap := make(map[int]int)
	countOfA := 0
	for i:= 0; i < len(substring); i++ {
		if substring[i] == 'a' {
			countOfA += 1
		}
		aMap[i] = countOfA
	}

	return aMap
}
/*
This is kahwei's logic with jynn's implementation:
1. Declare a map with []int, int
    - key value indicates the index of string,s given
    - value is to indicate the total number of string until the index(key) value
2. Calculate how many times need to repeat the string
    - times = divide the n with length of string
    - total number of a = last value in map x times
3. Calculate the remainder of how many characters left need to consider
    - total number of a = (last value in map x times) + map[remainder]
*/

func mapSolution(s string, n int64) int64 {
	//declare an empty map with type []int, int
	mapOfS := make (map[int] int64)
	//declare a counter to calculate the total number of char 'a'
	var totalOfA int64 = 0
	//loop through the string check the number of char 'a'
	//add the key(index) and total value of a into the map
	for i := 0; i < len(s); i++{
		if s[i] == 'a'{
			totalOfA ++
		}
		mapOfS[i] = totalOfA
	}
	//calculate how many times should repeat the string
	times := n / int64(len(s))
	//calculate the remainder of how many characters left that needed to consider
	remainder := n % int64(len(s))
	if remainder != 0{
		totalOfA = (totalOfA * times) + mapOfS[int(remainder-1)]
	}else{
		totalOfA = totalOfA * times
	}
	return totalOfA
}

/* This is Jynn's logic:
1. Split the string,s and append into an array of byte
    - get the length of the array / number of substring, len(s)
    - count the number of a, countOfA through a for loop
2. Calculate the how times need to repeat the string,s
    - by dividing the integer with the number of substring
    - n/len(s)
3. Calculate the modulus of n and len(s)
    - if modulus equals to 0 - countOfA * (n/len(s))
    - if modulus not equals to 0 - (countOfA * (n/len(s)))+1
4. Return countOfA
*/
func repeatedString(s string, n int64) int64 {
	//declare a counter to count the existence of "a"
	var countOfA int64
	//declare an empty slice with byte
	var characterArray []byte
	//split and append the substring into the array
	characterArray = append (characterArray, s...)
	//to count the number of a
	for _, v := range characterArray{
		if v == 'a'{
			countOfA++
		}
	}
	if countOfA != 0 {
		//to count how many times need to repeat the string
		timesToRepeat := n / int64(len(s))
		remainder := n % int64(len(s))

		if remainder == 0 {
			countOfA = countOfA * timesToRepeat
		}else{
			countOfA = (countOfA * timesToRepeat) + 1
		}

		return countOfA
	}else{
		return 0
	}
}