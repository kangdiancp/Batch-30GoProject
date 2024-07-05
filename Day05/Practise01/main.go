package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{2, 3, 5, 98, 76, 23}
	fmt.Println("Target ada di index : ", linearSearch(numbers, 98))

	//fmt.Println("Binary Search ada di index : ", binarySearch(numbers, 23))

	fmt.Println(numbers)

	//sumTarget(numbers, 28)
	fmt.Println("sumTarget : ", sumTarget(numbers, 28))

	fmt.Println("sumTargetOpt : ", sumTargetOptimize(numbers, 28))

	//binary search for string
	words := []string{"A", "B", "Z", "C", "X", "Y", "Z"}
	fmt.Println("Binary Search ada di index : ", binarySearchString(words, "C"))
	fmt.Println(words)
}

// 3. linear search/ brute force
func sumTarget(nums []int, target int) []int {
	//store return value
	sumIndex := []int{}

	// O(n * n)=O(n2)
	for i := 0; i < len(nums); i++ { //loop 1 -> O(n)
		for j := i + 1; j < len(nums); j++ { //loop 2 O(n), bottle neck
			if nums[i]+nums[j] == target { // O(1),O(2)..
				sumIndex = append(sumIndex, i, j)
			}
		}
	}
	return sumIndex
}

// 4. optimize
func sumTargetOptimize(nums []int, target int) []int {
	//1. declare map
	mapNums := make(map[int]int)
	//2. store return value
	sumIndex := []int{}

	//3. insert data from slice nums to mapNums// O(1)
	for index, v := range nums {
		mapNums[v] = index
	}

	for i, _ := range nums { // O(n)
		subtract := target - nums[i]

		//check is substract ada di mapNums ?
		if value, ok := mapNums[subtract]; ok { //O(1)
			sumIndex = append(sumIndex, value)
		}
	}

	return sumIndex

}

// 1. linear search
func linearSearch(nums []int, target int) int {
	for index, v := range nums {
		if v == target {
			return index
		}
	}
	return 0
}

// 2. binary search
func binarySearch(nums []int, target int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)
	for {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}

		if left > right {
			break
		}
	}
	return -1

}

// 2. binary search
func binarySearchString(words []string, target string) int {
	sort.Strings(words)
	left, right := 0, len(words)
	for {
		mid := (left + right) / 2

		if words[mid] == target {
			return mid
		} else if target < words[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}

		if left > right {
			break
		}
	}
	return -1

}
