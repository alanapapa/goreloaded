package main

import (
	"fmt"
)

func main() {
	arg1 := 4
	arg2 := 3
	fmt.Println(RecursivePower(arg1, arg2))
}

func RecursivePower(nb int, power int) int {
	res := 0
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		res = nb * RecursivePower(nb, power-1)
	}
	return res
}
