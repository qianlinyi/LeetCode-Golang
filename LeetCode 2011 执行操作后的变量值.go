package main

import "strings"

func finalValueAfterOperations(operations []string) int {
	ans := 0
	for _, operation := range operations {
		if strings.HasPrefix(operation, "--") || strings.HasSuffix(operation, "--") {
			ans--
		}
		if strings.HasPrefix(operation, "++") || strings.HasSuffix(operation, "++") {
			ans++
		}
	}
	return ans
}

func main() {}
