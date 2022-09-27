package main

type void struct {
}

func checkIfPangram(sentence string) bool {
	set := make(map[int32]void)
	var exist void
	for _, i := range sentence {
		set[i] = exist
	}
	return len(set) == 26
}
