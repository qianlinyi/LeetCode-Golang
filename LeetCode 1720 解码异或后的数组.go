package main

func decode(encoded []int, first int) []int {
	ans := make([]int, len(encoded)+1)
	ans[0] = first
	for i, num := range encoded {
		ans[i+1] = num ^ ans[i]
	}
	return ans
}

func main() {}
