package main

import "sort"

func twoSum(nums []int, target int) []int {

	for i, n := range nums {
		for j, m := range nums {

			if i == n {
				continue
			}

			if n+m == target {
				result := []int{i, j}
				return result
			}
		}
	}
	return nil
}

func twoSum_hash1(nums []int, target int) []int {

	hashTable := make(map[int]int, len(nums))
	for i, n := range nums {
		hashTable[n] = i
	}

	for j, m := range nums {
		complement := target - m
		indice, ok := hashTable[complement]
		if ok && indice != j {
			return []int{j, indice}
		}
	}
	return nil
}

func twoSum_hash2(nums []int, target int) []int {

	hashTable := make(map[int]int, len(nums))
	for j, m := range nums {
		complement := target - m
		indice, ok := hashTable[complement]
		if ok {
			result := []int{j, indice}
			sort.Ints(result)
			return result
		}
		hashTable[m] = j
	}
	return nil
}
