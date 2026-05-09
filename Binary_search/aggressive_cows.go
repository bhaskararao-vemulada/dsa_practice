package main

import (
	"fmt"
	"sort"
)

func canPlace(stalls []int, cows int, distance int) bool {
	placedCows := 1
	lastPlaced := stalls[0]

	for i := 1; i < len(stalls); i++ {

		if stalls[i]-lastPlaced >= distance {
			placedCows++
			lastPlaced = stalls[i]
		}

		if placedCows == cows {
			return true
		}
	}

	return false
}

func aggressiveCows(stalls []int, cows int) int {
	sort.Ints(stalls)

	left := 1
	right := stalls[len(stalls)-1] - stalls[0]
	ans := -1

	for left <= right {
		mid := left + (right-left)/2

		if canPlace(stalls, cows, mid) {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

func main() {
	stalls := []int{1, 2, 4, 8, 9}
	cows := 3

	x := aggressiveCows(stalls, cows)
	fmt.Println(x)
}