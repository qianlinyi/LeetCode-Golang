package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subtractProductAndSum(n int) int {
	nums, sum, product := strings.Split(strconv.Itoa(n), ""), 0, 1
	for _, num := range nums {
		digit, _ := strconv.Atoi(num)
		sum, product = sum+digit, product*digit
	}
	return product - sum
}

func main() {
	n := 123
	fmt.Println(strings.Split(strconv.Itoa(n), ""))
}
