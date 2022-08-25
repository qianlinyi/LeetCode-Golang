package main

func sumOddLengthSubarrays(arr []int) (sum int) {
	n := len(arr)
	for i, v := range arr {
		leftCount, rightCount := i, n-i-1
		leftOdd := (leftCount + 1) / 2
		rightOdd := (rightCount + 1) / 2
		leftEven := leftCount/2 + 1
		rightEven := rightCount/2 + 1
		sum += v * (leftOdd*rightOdd + leftEven*rightEven)
	}
	return sum
}
