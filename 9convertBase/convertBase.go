package main

import "fmt"

func ConvertBase(nbr, baseFrom, baseTo string) string {
	goo := AtoiB(nbr, baseFrom)
	return NbrB(goo, baseTo)
}

func AtoiB(str string, base string) int {
	var count, at int
	for range base {
		count++
	}
	for _, r := range str {
		for i, k := range base {
			if r == k {
				at = at*count + i
			}
		}
	}
	return at
}
func NbrB(num int, base string) string {
	var count int
	var nbr string
	for range base {
		count++
	}
	for num > 0 {
		nbr = string(base[num%count]) + nbr
		num = num / count
	}
	return nbr
}
func main() {
	result := ConvertBase("4506C", "0123456789ABCDEF", "choumi")
	fmt.Println(result)
	result = ConvertBase("babcbaaaaabcb", "abc", "0123456789ABCDEF")
	fmt.Println(result)
	result = ConvertBase("256850", "0123456789", "01")
	fmt.Println(result)
	result = ConvertBase("uuhoumo", "choumi", "Zone01")
	fmt.Println(result)
	result = ConvertBase("683241", "0123456789", "01")
	fmt.Println(result)
}
