package main

import (
	"math/rand"
)

// WeightedChoice 根据权重随机，返回对应选项的索引，O(n)
func WeightedChoice(weightArray []int) int {
	if weightArray == nil {
		return -1
	}
	total := Sum(weightArray)
	if total <= 0 {
		return -1
	}
	rv := rand.Intn(total)
	for i, v := range weightArray {
		if rv < v {
			return i
		}
		rv -= v
	}
	return len(weightArray) - 1
}

// Sum 求和
func Sum(data []int) int {
	var sum int
	for _, v := range data {
		sum += v
	}
	return sum
}
