package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print(isPalindrome(1234564321))
}

func isPalindrome(x int) bool {
	temp := strconv.Itoa(x)
	counter := 0
	for i:= 0; i < len(temp);i++{
		if temp[i] == temp[len(temp) - i - 1] {
			counter++
		}else{
			return false
		}
		if len(temp)/2+1 == counter{
			return true
		}
	}
	return true
}