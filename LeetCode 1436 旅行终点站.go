package main

func destCity(paths [][]string) string {
	cntl, cntr := make(map[string]int), make(map[string]int)
	for _, path := range paths {
		cntl[path[0]]++
		cntr[path[1]]++
	}
	for key, value := range cntr {
		_, ok := cntl[key]
		if value == 1 && !ok {
			return key
		}
	}
	return ""
}
