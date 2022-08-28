package main

func anagramMappings(nums1 []int, nums2 []int) []int {
	pos, ans := make(map[int][]int), make([]int, len(nums1))
	for index, num := range nums2 {
		pos[num] = append(pos[num], index)
	}
	for index, num := range nums1 {
		ans[index] = pos[num][0]
		if len(pos[num]) > 1 {
			pos[num] = pos[num][1:]
		}
	}
	return ans
}
