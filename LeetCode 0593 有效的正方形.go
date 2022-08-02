package main

import "sort"

func distance(p1 []int, p2 []int) int {
	return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
}

func isIsoscelesTriangle(p1 []int, p2 []int, p3 []int) bool {
	var side []int
	side = append(side, distance(p1, p2))
	side = append(side, distance(p1, p3))
	side = append(side, distance(p2, p3))
	sort.Ints(side)
	return side[0] != 0 && side[0] == side[1] && side[0]+side[1] == side[2]
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	return isIsoscelesTriangle(p1, p2, p3) && isIsoscelesTriangle(p1, p2, p4) && isIsoscelesTriangle(p1, p3, p4) && isIsoscelesTriangle(p2, p3, p4)
}
