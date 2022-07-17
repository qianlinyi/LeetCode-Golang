package main

func minPartitions(n string) int {
	ans := 0
	cnt := 0
	for _, i := range n {
		if cnt < int(i-'0') {
			ans += int(i-'0') - cnt
			cnt += int(i-'0') - cnt
		}
	}
	return ans
}

func main() {}
