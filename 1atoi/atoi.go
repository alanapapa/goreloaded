package main

import (
	"fmt"
)

func main() {
	s := "12345"
	s2 := "0000000012345"
	s3 := "012 345"
	s4 := "Hello World!"
	s5 := "+1234"
	s6 := "-1234"
	s7 := "++1234"
	s8 := "--1234"

	n := Atoi(s)
	n2 := Atoi(s2)
	n3 := Atoi(s3)
	n4 := Atoi(s4)
	n5 := Atoi(s5)
	n6 := Atoi(s6)
	n7 := Atoi(s7)
	n8 := Atoi(s8)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)
	fmt.Println(n7)
	fmt.Println(n8)
}

func Atoi(s string) int {
	s0 := s
	num := 0

	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}

	for _, e := range s {
		if e >= 48 && e <= 57 {
			num = (int(e) - 48) + num*10
		} else {
			return 0
		}
	}

	if s0[0] == '-' {
		num = -num
	}
	return num
}
