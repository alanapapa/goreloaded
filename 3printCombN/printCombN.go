package main

import "github.com/01-edu/z01"

func main() {
	PrintCombN(1)
	PrintCombN(3)
	PrintCombN(9)
}

func PrintCombN(n int) {
	abc(0, n+2, n, ", ")
	z01.PrintRune('\n')
}

func abc(start, end, n int, out string) {
	base := ", 012345678"
	if n == 0 {
		if out == base[:end] {
			for _, e := range out {
				if e >= '0' && e <= '9' {
					z01.PrintRune(e)
				}
			}
			return
		}
		for _, e := range out {
			z01.PrintRune(e)
		}
	}

	for i := start; i <= 9; i++ {
		str := out + string(i+48)
		abc(i+1, end, n-1, str)
	}

}
