package main

func canConstruct(ransomNote string, magazine string) bool {
	cnt1, cnt2 := make([]int, 130), make([]int, 130)
	for _, c := range ransomNote {
		cnt1[c]++
	}
	for _, c := range magazine {
		cnt2[c]++
	}
	flag := true
	for i := 0; i < 130; i++ {
		if cnt1[i] > cnt2[i] {
			flag = false
			break
		}
	}
	return flag
}
