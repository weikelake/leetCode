package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(romanToInt("DCLXXIX"))
}

func romanToInt(s string) int {
	doubleSymbols := map[string]int{
		"IV" : 4,
		"IX" : 9,
		"XL" : 40,
		"XC" : 90,
		"CD" : 400,
		"CM" : 900,
	}

	symbols := map[string]int{
		"I" : 1,
		"V" : 5,
		"X" : 10,
		"L" : 50,
		"C" : 100,
		"D" : 500,
		"M" : 1000,
	}

	result := 0

	for k, value := range doubleSymbols {
		if strings.Count(s,k) > 0{
			result += value * strings.Count(s,k)
			s = strings.Replace(s, k, "", -1)
		}
	}

	for i:=0; i<len(s); i++{
		result+=symbols[string(s[i])]
	}

	return result
}