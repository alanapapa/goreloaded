package main

import "fmt"

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			first := array[i]
			second := array[j]
			if f(first, second) == 1 {
				array[i], array[j] = array[j], array[i]
			} else if f(first, second) == -1 {
			} else if f(first, second) == 0 {
			}
		}
	}
}

func Compare(a, b string) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}
func RevCompare(b, a string) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}
func main() {
	result := []string{"The", "earliest", "computing", "device", "undoubtedly", "consisted", "of", "the", "five", "fingers", "of", "each", "hand"}
	AdvancedSortWordArr(result, Compare)
	fmt.Println(result)
	result = []string{"The", "word", "digital", "comesfrom", "\"digits\"", "or", "fingers"}
	AdvancedSortWordArr(result, Compare)
	fmt.Println(result)
	result = []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	AdvancedSortWordArr(result, Compare)
	fmt.Println(result)
	result = []string{"The", "computing", "consisted", "device", "each", "earliest", "fingers", "five", "hand", "of", "of", "the", "undoubtedly"}
	AdvancedSortWordArr(result, RevCompare)
	fmt.Println(result)
	result = []string{"\"digits\"", "The", "comesfrom", "digital", "fingers", "or", "word"}
	AdvancedSortWordArr(result, RevCompare)
	fmt.Println(result)
	result = []string{"1", "2", "3", "A", "B", "C", "a", "b", "c"}
	AdvancedSortWordArr(result, RevCompare)
	fmt.Println(result)
}
