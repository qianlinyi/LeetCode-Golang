package main

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			break
		} else {
			digits[i] = 0
		}
	}
	if digits[0] == 0 {
		digits = append(digits, 0)
		copy(digits[1:], digits)
		digits[0] = 1
	}
	return digits
}

func main() {}
