package main

//hash表解法
func twoSum(nums []int, target int) []int {
	sum := make(map[int]int)
	for k,v :=range nums {
		if k1,ok := sum[v]; ok {
			return []int{k1,k}
		}
		sum[target-v] = k
	}
	return []int{}
}