package main

import (
	"fmt"
	"strings"
)

func Split(str, charset string) []string {
	l1 := len(str)
	l2 := len(charset)
	size := 0
	for i := 0; i < l1-l2; i++ {
		if str[i:i+l2] == charset {
			size++
		}
	}
	s := make([]string, size+1)
	j := 0
	mount := 0
	for i := 0; i <= l1-l2; i++ {
		if str[i:i+l2] == charset {
			s[j] = str[mount:i]
			j++
			mount = l2 + i
		} else if i == l1-l2 {
			s[j] = str[mount:]
		}
	}
	return s
}

func main() {

	str := "|=choumi=|which|=choumi=|itself|=choumi=|used|=choumi=|cards|=choumi=|and|=choumi=|a|=choumi=|central|=choumi=|computing|=choumi=|unit.|=choumi=|When|=choumi=|the|=choumi=|machine|=choumi=|was|=choumi=|finished,"
	fmt.Println(Split(str, "|=choumi=|"))
	str = "AFJCharlesAFJBabbageAFJstartedAFJtheAFJdesignAFJofAFJtheAFJfirstAFJautomaticAFJmechanicalAFJcalculator,"
	fmt.Println(Split(str, "AFJ"))
	str = "<<==123==>>In<<==123==>>1820,<<==123==>>Thomas<<==123==>>de<<==123==>>Colmar<<==123==>>launched<<==123==>>the<<==123==>>mechanical<<==123==>>calculator<<==123==>>industry[note<<==123==>>1]<<==123==>>when<<==123==>>he<<==123==>>released<<==123==>>his<<==123==>>simplified<<==123==>>arithmometer,"
	fmt.Println(Split(str, "<<==123==>>"))

	fmt.Println(strings.Split("Hello", ""))
	fmt.Println(Split("Hello", ""))
}
