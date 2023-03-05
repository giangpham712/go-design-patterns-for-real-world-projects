package main

import (
	"math/rand"
	"time"
)

type numbers [1024 * 1024]int

func main() {
	var nums *numbers = new(numbers)
	initialize(nums)
}

func initialize(nums *numbers) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(nums); i++ {
		nums[i] = rand.Intn(10000)
	}
}

func max(nums *numbers) int {
	temp := nums[0]
	for _, val := range nums {
		if val > temp {
			temp = val
		}
	}
	return temp
}
