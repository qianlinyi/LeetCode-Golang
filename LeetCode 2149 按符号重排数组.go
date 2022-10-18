package main

func rearrangeArray(nums []int) []int {
	ans, q1, q2 := make([]int, 0), make([]int, 0), make([]int, 0)
	for _, num := range nums {
		if num > 0 {
			q1 = append(q1, num)
		} else {
			q2 = append(q2, num)
		}
	}
	for i := 0; i < len(q1); i++ {
		ans = append(ans, q1[i])
		ans = append(ans, q2[i])
	}
	return ans
}
