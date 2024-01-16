package main

/*
维护两张哈希表，第一张哈希表s2t以s中字符为键，映射至t的字符为值，
第二张哈希表t2s以t中字符为键，映射至s的字符为值。
从左至右遍历两个字符串的字符，不断更新两张哈希表，如果出现冲突
（即当前下标index对应的字符s[index]已经存在映射且不为t[index]或
当前下标index对应的字符t[index]已经存在映射且不为s[index]时说明两个字符串无法构成同构，返回false。
*/
func isIsomorphic(s string, t string) bool {
	s2t, t2s := make(map[byte]byte), make(map[byte]byte)
	for i := range s {
		x, y := s[i], t[i]
		if s2t[x] > 0 && s2t[x] != y || t2s[y] > 0 && t2s[y] != x {
			return false
		}
		s2t[x] = y
		t2s[y] = x
	}
	return true
}
