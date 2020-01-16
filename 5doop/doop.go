package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	count := 0
	res := "0"
	for range args {
		count++
	}
	if count != 3 {
		res = ""
		RangePrint(res)
	}
	if Valid(args) && MaxMin(args[0]) && MaxMin(args[2]) {
		res = Doop(args)
		RangePrint(res)
		z01.PrintRune(10)
	} else {
		RangePrint(res)
		z01.PrintRune(10)
	}
}
func RangePrint(str string) {
	for _, r := range str {
		z01.PrintRune(r)
	}
}
func Valid(arr []string) bool {
	for _, r := range arr[0] {
		if (r < '0' || r > '9') && r != '-' {
			return false
		}
	}
	for _, k := range arr[1] {
		if k != '-' && k != '+' && k != '/' && k != '*' && k != '%' {
			return false
		}
	}
	for _, b := range arr[2] {
		if (b < '0' || b > '9') && b != '-' {
			return false
		}
	}
	return true
}
func MaxMin(str string) bool {
	max := "9223372036854775807"
	min := "-9223372036854775808"
	if str[0] == '-' {
		if NeLen(str) > NeLen(min) {
			return false
		}
		if NeLen(str) == NeLen(min) {
			if str > min {
				return false
			}
		}
	} else {
		if NeLen(str) > NeLen(max) {
			return false
		}
		if NeLen(str) > NeLen(max) {
			if str > max {
				return false
			}
		}
	}
	return true
}
func NeLen(str string) int {
	count := 0
	for range str {
		count++
	}
	return count
}
func Doop(arr []string) string {
	a := Atoi(arr[0])
	b := Atoi(arr[2])
	x := arr[1]
	res := 0
	if x == "-" {
		res = a - b
		if (a > 0 && b < 0 && res < 0) || (a < 0 && b > 0 && res > 0) {
			return "0"
		}
	}
	if x == "+" {
		res = a + b
		if (a > 0 && b > 0 && res < 0) || (a < 0 && b < 0 && res > 0) {
			return "0"
		}
	}
	if x == "*" {
		res = a * b
		if res/a != b {
			return "0"
		}
	}
	if x == "/" {
		if b == 0 {
			return "No division by 0"
		}
		res = a / b

	}
	if x == "%" {
		if b == 0 {
			return "No modulo by 0"
		}
		res = a % b
	}
	return Itoa(res)
}
func Atoi(str string) int {
	str0 := str
	var num int
	if str[0] == '-' || str[0] == '+' {
		str = str[1:]
	}
	for _, r := range str {
		if r >= '0' && r <= '9' {
			num = int((r)-48) + num*10
		} else {
			return 1
		}
	}
	if str0[0] == '-' {
		num = -num
	}
	return num
}
func Itoa(num int) string {
	num0 := num
	var str string
	if num == 0 {
		str = "0"
	}
	if num == -9223372036854775808 {
		str = "9223372036854775808"
	}
	if num < 0 {
		num = num * -1
	}
	for num > 0 {
		str = string(num%10+48) + str
		num = num / 10
	}
	if num0 < 0 {
		str = "-" + str
	}
	return str
}

// ./5doop 861 + 870
// ./5doop 861 - 870
// ./5doop 861 "*" 870
// ./5doop 861 % 870
// ./5doop hello + 1
// ./5doop 1 p 1
// ./5doop 1 / 0
// ./5doop 1 % 0
// ./5doop 1 "*" 1
// ./5doop 9223372036854775807 + 1
// ./5doop 9223372036854775809 - 3
// ./5doop 9223372036854775807 "*" 3
