package benchmark

import (
	"day06/practise02/quiz"
	"testing"
)

//var nums = []int{4, 2, 3, 11, 15, 19}

var nums = quiz.CreteIntSlice(1_000_000, 100_000)
var target = 55_555

func BenchmarkAllTwoSum(b *testing.B) {
	b.Run("twoSumBruteForce", twoSumBruteForce)
	b.Run("twoSumOptimize", twoSumOptimize)
}

func twoSumBruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		quiz.TwoSumBruteForce(nums, target)
	}
}

func twoSumOptimize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		quiz.TwoSumOptimize(nums, target)
	}
}
