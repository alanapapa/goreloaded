package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	count := 0
	for range args {
		count++
	}
	if count == 0 {
		for {
			buff := []byte{0}
			os.Stdin.Read(buff)
			for _, r := range buff {
				z01.PrintRune(rune(r))
			}
		}
	} else {
		for _, r := range args {
			read, err := ioutil.ReadFile(r)
			if err == nil {
				for _, x := range read {
					z01.PrintRune(rune(x))
				}
			} else {
				for _, m := range err.Error() {
					z01.PrintRune(m)
				}
			}
		}
		z01.PrintRune(10)
	}
}

// ./12cat
// Hello
// jaflsdfj
// Computer science (sometimes called computation science or computing science

// ./12cat quest8.txt
// ./12cat quest8T.txt
// ./12cat quest8T.txt quest8.txt
