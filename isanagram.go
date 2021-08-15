package main

//hash表解法
func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	mapS := make(map[rune]int)
	for _,v:=range s {
		mapS[v]++
	}
	for _,v := range t {
		mapS[v]--
		if mapS[v] < 0 {
			//如果不相等 t中必然有至少一个字符的结果被减成负的
			return false
		}
	}
	return true
}