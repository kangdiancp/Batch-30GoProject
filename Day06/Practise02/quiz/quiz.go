package quiz

import (
	"crypto/rand"
	"math/big"
)

func AddSum(a, b int) int {
	return a + b
}

// 1. //brute force solution/linear search
func TwoSumBruteForce(numbers []int, target int) []int {
	sumIndex := []int{}
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				sumIndex = append(sumIndex, i, j)
			}
		}

	}
	return sumIndex
}

// 2. using map
func TwoSumOptimize(numbers []int, target int) []int {
	maps := make(map[int]int)
	sumIndex := []int{}
	// insert into map
	for i, v := range numbers {
		maps[v] = i
	}

	for i, _ := range numbers {
		b := target - numbers[i]

		if value, ok := maps[b]; ok {
			sumIndex = append(sumIndex, value)
		}
	}

	return sumIndex
}

func CreteIntSlice(length int, maxValue int) []int {
	var result = make([]int, 0, length)
	reader := rand.Reader
	for i := 0; i < length; i++ {
		random, _ := rand.Int(reader, big.NewInt(int64(maxValue)))
		result = append(result, int(random.Int64()))
	}
	return result
}
