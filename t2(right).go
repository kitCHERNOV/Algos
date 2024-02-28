// Leetcode 643 Maximum average subarray 1

//
package main

import (
	"fmt"
	"math"
)

// All vars and constants:
// var nums = []int{1,12,-5,-6,50,3}
// var n = len(nums)

// 1 // using algo of prefix subarray

// func findMaxaverage(nums []int, k int) (res float64){
// 	// Создаем массив префексных сумм
// 	var sums []int = make([]int, n)
// 	// Инициализация первого элементы тк. можем выйти за пределы
// 	sums[0] = nums[0]
// 	for i:=1 /* 1 потомучто нулевой эл-т уже объявлен*/; i < len(sums); i++ {
// 		sums[i] = sums[i-1] + nums[i]
// 	}

// 	res = float64(sums[k-1]) / float64(k)
// 	for i:= k; i < len(sums); i++ {
// 		res = math.Max(res, float64(sums[i] - sums[i-k]) / float64(k))
// 	} 
// 	return
// }

// 2 using method of sliding window

func findMaxaverage(nums []int, k int) float64 {
	sum := 0.
	res := 0.

	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}

	for i := k; i < len(nums); i++ {
		sum += float64(nums[i] - nums[i-k])
		res = math.Max(res, sum)
	}
	
	return res / float64(k)
}

func main() {
	var nums = []int{1,12,-5,-6,50,3}
	//var n = len(nums)
	k := 4
	fmt.Println(findMaxaverage(nums, k))
}