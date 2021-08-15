package main

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)
	for _,v:= range strs  {
		strS := []byte(v)
		sort.Slice(strS, func(i,j int)bool {return strS[i]< strS[j]})
		str := string(strS)
		strMap[str] = append(strMap[str],v)
	}
	res := make([][]string,0,len(strMap))
	for _,v :=range strMap {
		res = append(res, v)
	}
	return res
}