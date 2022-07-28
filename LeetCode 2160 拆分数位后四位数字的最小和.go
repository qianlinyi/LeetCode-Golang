package main

import (
	"sort"
)

func minimumSum(num int) int {
	var digit []int
	for num > 0 {
		digit = append(digit, num%10)
		num /= 10
	}
	sort.Ints(digit)
	return (digit[0]+digit[1])*10 + (digit[2] + digit[3])
}

func main() {
}
