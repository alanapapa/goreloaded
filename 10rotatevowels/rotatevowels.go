package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	var count, sound int
	var str string
	for range args {
		count++
	}
	for _, r := range args {
		for _, x := range r {
			if Vowels(x) {
				str = str + string(x)
				sound++
			}
		}
	}
	for ind, r := range args {
		for _, w := range r {
			if Vowels(w) {
				z01.PrintRune(rune(str[sound-1]))
				sound--
			} else {
				z01.PrintRune(w)
			}
		}
		if ind < count-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune(10)
}
func Vowels(run rune) bool {
	vowels := "AEIOUaeiou"
	for _, r := range vowels {
		if r == run {
			return true
		}
	}
	return false
}
