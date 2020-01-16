package main

import "fmt"

func SplitWhiteSpaces(str string) []string {
	var space rune = ' '
	var count, mount int
	var word string
	for _, r := range str {
		if r != 32 && r != 9 && r != 10 && (space == 32 || space == 9 || space == 10) {
			count++
		}
		space = r
	}
	arrgh := make([]string, count)
	for _, s := range str {
		if s != 32 && s != 9 && s != 10 {
			word = word + string(s)
		} else if (s == 32 || s == 9 || s == 10) && word != "" {
			arrgh[mount] = word
			mount++
			word = ""

		}
	}
	if word != "" {
		arrgh[mount] = word
	}
	return arrgh
}

func main() {
	str := "The earliest foundations of what would become computer science predate the invention of the modern digital computer"
	fmt.Println(SplitWhiteSpaces(str))
	str = "    Machines for calculating fixed numerical tasks such as the abacus have existed since antiquity,     "
	fmt.Println(SplitWhiteSpaces(str))
	str = "aiding in computations such as multiplication and division ."
	fmt.Println(SplitWhiteSpaces(str))
	str = " Algorithms for performing computations have existed since antiquity, even before the development of sophisticated computing equipment."
	fmt.Println(SplitWhiteSpaces(str))
	str = "Wilhelm Schickard designed and constructed the first working mechanical calculator in 1623.[4]"
	fmt.Println(SplitWhiteSpaces(str))
	str = "In 1673, Gottfried Leibniz demonstrated a digital mechanical calculator,"
	fmt.Println(SplitWhiteSpaces(str))
}
