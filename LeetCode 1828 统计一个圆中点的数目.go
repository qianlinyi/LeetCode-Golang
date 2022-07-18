package main

func countPoints(points [][]int, queries [][]int) []int {
	ans := make([]int, len(queries))
	for index, query := range queries {
		cnt := 0
		for _, point := range points {
			if (point[0]-query[0])*(point[0]-query[0])+(point[1]-query[1])*(point[1]-query[1]) <= query[2]*query[2] {
				cnt++
			}
		}
		ans[index] = cnt
	}
	return ans
}

func main() {}
