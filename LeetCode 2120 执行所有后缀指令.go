package main

func move(pos []int, c int32) []int {
	if c == 'R' {
		return []int{pos[0], pos[1] + 1}
	} else if c == 'L' {
		return []int{pos[0], pos[1] - 1}
	} else if c == 'U' {
		return []int{pos[0] - 1, pos[1]}
	} else if c == 'D' {
		return []int{pos[0] + 1, pos[1]}
	}
	return []int{}
}

func judge(pos []int, n int) bool {
	if pos[0] < 0 || pos[0] >= n || pos[1] < 0 || pos[1] >= n {
		return false
	} else {
		return true
	}
}

func executeInstructions(n int, startPos []int, s string) []int {
	ans := make([]int, len(s))
	for i, _ := range s {
		pos := startPos
		ans[i] = len(s) - i
		for j, c := range s[i:] {
			pos = move(pos, c)
			if judge(pos, n) == false {
				ans[i] = j
				break
			}
		}
	}
	return ans
}
