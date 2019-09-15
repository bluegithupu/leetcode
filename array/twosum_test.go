package main

import (
	"log"
	"testing"
)

var (
	nums   = []int{2, 7, 11, 15}
	target = 9
)

func Test_twoSum(t *testing.T) {
	r := twoSum(nums, target)

	if r[0] == 0 && r[1] == 1 {
		t.Log("yes")
	} else {
		t.Fatal("no")
	}
}

func Test_twoSum_hash1(t *testing.T) {
	r := twoSum_hash1(nums, target)

	if r[0] == 0 && r[1] == 1 {
		t.Log("yes")
	} else {
		t.Fatal("no")
	}
}

func Test_twoSum_hash2(t *testing.T) {
	r := twoSum_hash2(nums, target)
	log.Println(r)

	if r[0] == 0 && r[1] == 1 {
		t.Log("yes")
	} else {
		t.Fatal("no")
	}
}
