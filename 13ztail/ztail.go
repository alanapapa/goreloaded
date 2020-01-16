package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	input := os.Args[1:]
	name := ""
	leng := 0
	for range input {
		leng++
	}
	var num int

	var skip []int
	GlobalSign := false
	for argNum := range input {
		name = input[argNum]
		_, sign, valid := Atoi(name)
		if input[argNum] == "-c" {
			skip = append(skip, argNum)
			if argNum == leng-1 {
				fmt.Printf("tail: option requires an argument -- 'c' \n")
				os.Exit(0)
				return
			} else {
				name = input[argNum+1]
				_, sign, valid = Atoi(name)
				if valid {
					fmt.Printf("tail: invalid number of bytes: '%s' \n", name)
					os.Exit(0)
				} else {
					skip = append(skip, argNum+1)
					if sign {
						GlobalSign = sign
					}
				}
			}
			continue
		}

		if valid1, valid2 := CheckC(name); valid1 {
			skip = append(skip, argNum)
			if valid2 {
				fmt.Printf("tail: invalid number of bytes: '%s' \n", name[2:])
				return
			}
			if _, sign, _ = Atoi(name[2:]); sign {
				GlobalSign = sign
			}
			continue
		}
	}
	for i := ArrLen(skip); i > 0; i-- {
		name = input[skip[i-1]]
		_, _, valid := Atoi(name)
		valid1, _ := CheckC(name)
		if valid1 {
			num, _, _ = Atoi(name[2:])
			break
		}
		if !valid {
			num, _, _ = Atoi(name)
			break
		}
	}
	if num == 0 && !GlobalSign {
		return
	}
	for argNum := range input {
		name = input[argNum]
		if SkipCheck(skip, argNum) {
			continue
		}
		file, err := os.Open(name)
		if err != nil {
			fmt.Printf("tail: cannot open '%s' for reading: No such file or directory \n", name)

		} else {

			fileInfo, err := file.Stat()
			if err != nil {
				log.Fatal(err)
			}
			a := int(fileInfo.Size())
			arr := []byte{}
			for i := 0; i < a; i++ {
				arr = append(arr, 0)
			}
			file.Read(arr)
			if leng-ArrLen(skip) != 1 {
				fmt.Printf("==> %s <== \n", name)
			}
			if num > a {
				num = a
			}

			if GlobalSign {
				if num == 0 {
					fmt.Printf(string(arr[:]))
				} else {
					fmt.Printf(string(arr[num-1:]))
				}

			} else {

				fmt.Printf(string(arr[a-num:]))
			}
			if leng-2 != argNum+1 && leng != argNum+1 {
				fmt.Printf("\n")
			}

		}
	}
	os.Exit(0)

}

func Atoi(word string) (int, bool, bool) {
	isPlus := false
	invalid := false
	tempWord := word

	if word[0] == '+' {
		isPlus = true
		tempWord = word[1:]
	}
	if isPlus && tempWord == "" {
		invalid = true
	}
	if word[0] == '-' {
		tempWord = word[1:]
	}
	converted := 0
	for _, val := range tempWord {
		val = val - '0'
		if val > 9 {
			invalid = true
			return 0, isPlus, invalid
		}
		converted = converted*10 + int(val)
	}
	return converted, isPlus, invalid
}
func CheckC(str string) (bool, bool) {
	length := 0
	for range str {
		length++
	}
	if length > 2 && str[:2] == "-c" {
		_, _, check := Atoi(str[2:])
		if !check {
			return true, false
		} else {
			return true, true
		}
	}
	return false, false
}

func SkipCheck(n []int, x int) bool {
	for _, i := range n {
		if x == i {
			return true
		}
	}
	return false
}
func ArrLen(n []int) int {
	i := 0
	for range n {
		i++
	}
	return i
}

//./13ztail -c 20 test.txt
//tail -c 20 test.txt
//./13ztail test.txt -c 23
//tail test.txt -c 23
//./13ztail -c jelrjq 15
//tail -c jelrjq 15
//./13ztail -c 11 test.txt jfdklsa
//tail -c 11 test.txt jfdklsa
//./13ztail 11 -c test.txt
//tail 11 -c test.txt
