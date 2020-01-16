package main

import "fmt"

func ActiveBits(n int) uint {
	var str string
	count := 0
	n0 := n
	if n < 0 {
		n *= -1
	}
	for n > 0 {
		str = string(n%2) + str
		n /= 2
	}
	for _, r := range str {
		if r == 1 {
			count++
		}
	}
	if n0 < 0 {
		count++
	}
	return uint(count)
}
func main() {
	nbits := ActiveBits(15)
	fmt.Println(nbits)
	nbits = ActiveBits(-7)
	fmt.Println(nbits)
	nbits = ActiveBits(4)
	fmt.Println(nbits)
	nbits = ActiveBits(11)
	fmt.Println(nbits)
	nbits = ActiveBits(9)
	fmt.Println(nbits)
	nbits = ActiveBits(12)
	fmt.Println(nbits)
	nbits = ActiveBits(2)
	fmt.Println(nbits)
}
