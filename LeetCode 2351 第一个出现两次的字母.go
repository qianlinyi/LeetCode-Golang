package main

func repeatedCharacter(s string) byte {
	cnt := make([]int, 26)
	for _, c := range s {
		cnt[c-97]++
		if cnt[c-97] == 2 {
			return byte(c)
		}
	}
	return 0
}

func main() {}
