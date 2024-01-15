package main

import (
	"fmt"
	"sort"
)

func kmpSearchAll(text, pattern string) []int {
	buildPartialMatchTable := func(pattern string) []int {
		table := make([]int, len(pattern))
		j := 0
		for i := 1; i < len(pattern); i++ {
			for j > 0 && pattern[i] != pattern[j] {
				j = table[j-1]
			}
			if pattern[i] == pattern[j] {
				j++
			}
			table[i] = j
		}
		return table
	}

	table := buildPartialMatchTable(pattern)
	i, j := 0, 0
	positions := []int{}
	for i < len(text) {
		if text[i] == pattern[j] {
			i++
			j++
			if j == len(pattern) {
				positions = append(positions, i-j)
				j = table[j-1]
			}
		} else {
			if j > 0 {
				j = table[j-1]
			} else {
				i++
			}
		}
	}

	return positions
}

func beautifulIndices(s string, a string, b string, k int) []int {
	posa, posb, ans := make([]int, 0), make([]int, 0), make([]int, 0)
	cnta, cntb, cnt := make([]int, 26), make([]int, 26), make([][]int, 26)
	for i := range cnt {
		cnt[i] = make([]int, len(s))
	}
	n, na, nb := len(s), len(a), len(b)
	for _, c := range a {
		cnta[int(c)-97]++
	}
	for _, c := range b {
		cntb[int(c)-97]++
	}
	for i, c := range s {
		cnt[int(c)-97][i]++
		if i > 0 {
			for j := 0; j < 26; j++ {
				cnt[j][i] += cnt[j][i-1]
			}
		}
	}
	//fmt.Println(cnta, cntb, cnt)
	for i := 0; i < n; i++ {
		if i <= n-na {
			cnt_tmp := make([]int, 26)
			flag := true
			for j := 0; j < 26; j++ {
				if i == 0 {
					cnt_tmp[j] = cnt[j][i+na-1]
				} else {
					cnt_tmp[j] = cnt[j][i+na-1] - cnt[j][i-1]
				}
				if cnta[j] != cnt_tmp[j] {
					flag = false
					break
				}
			}
			if flag {
				posa = append(posa, i)
			}
		}
		if i <= n-nb {
			cnt_tmp := make([]int, 26)
			flag := true
			for j := 0; j < 26; j++ {
				if i == 0 {
					cnt_tmp[j] = cnt[j][i+nb-1]
				} else {
					cnt_tmp[j] = cnt[j][i+nb-1] - cnt[j][i-1]
				}
				if cntb[j] != cnt_tmp[j] {
					flag = false
					break
				}
			}
			if flag {
				posb = append(posb, i)
			}
		}
	}
	fmt.Println(posa, posb)
	possa, possb := kmpSearchAll(s, a), kmpSearchAll(s, b)
	fmt.Println(possa, possb)
	if len(posa) == 0 || len(posb) == 0 {
		return ans
	}
	for _, pos := range posa {
		p := sort.Search(len(posb), func(i int) bool {
			return posb[i] >= pos
		})
		//fmt.Println(p)
		if p == len(posb) {
			if pos-posb[len(posb)-1] <= k {
				ans = append(ans, pos)
			}
		} else {
			if posb[p]-pos <= k || (p > 0 && pos-posb[p-1] <= k) {
				ans = append(ans, pos)
			}
		}
	}
	return ans
}
