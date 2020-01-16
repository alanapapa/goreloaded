package main

import "github.com/01-edu/z01"

func main() {
	PrintNbrBase(125, "0123456789")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
}

func PrintNbrBase(nbr int, base string) {
	res := ""
	num := nbr
	lenBase := 0
	for range base {
		lenBase++
	}
	for i := 0; i < lenBase; i++ {
		for j := i + 1; j < lenBase; j++ {
			if base[i] == base[j] || base[i] == '-' || base[i] == '+' {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}
	for nbr != 0 {
		index := nbr % lenBase
		if index < 0 {
			index *= -1
		}
		res = string(base[index]) + res
		nbr /= lenBase
	}
	if num < 0 {
		z01.PrintRune('-')
	}
	for _, e := range res {
		z01.PrintRune(e)
	}
}
