package main

import "fmt"

func AtoiBase(s string, base string) int {
	var res int
	if Check(base) {
		for _, first := range s {
			for i, e := range base {
				if first == e {
					res = res*BaseLen(base) + i
				}
			}
		}
	} else {
		return 0
	}
	return res
}

func Check(str string) bool {
	for i := 0; i < BaseLen(str); i++ {
		for j := i + 1; j < BaseLen(str); j++ {
			if str[i] == str[j] || str[i] == 43 || str[i] == 45 || BaseLen(str) < 2 {
				return false
			}
		}
	}
	return true
}
func BaseLen(str string) int {
	count := 0
	for range str {
		count++
	}
	return count
}
func main() {
	fmt.Println(AtoiBase("bcbbbbaab", "abc"))
	fmt.Println(AtoiBase("0001", "01"))
	fmt.Println(AtoiBase("00", "01"))
	fmt.Println(AtoiBase("saDt!I!sI", "CHOUMIisDAcat!"))
	fmt.Println(AtoiBase("AAho?Ao", "WhoAmI?"))
	fmt.Println(AtoiBase("thisinputshouldnotmatter", "abca"))
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}
