package main

import (
	"fmt"
)

// 2d task - Вторая задача — Максимальное среднее значение
// Дан целочисленный массив nums, состоящий из n элементов, и целое число k.
// Найдите непрерывный подмассив длиной k, имеющий максимальное среднее значение, и верните это значение.
// Подмассив — это последовательность чисел без изменения порядка или пропуска элементов исходного массива.

func Addition(slice []int) (sum float64) { // de facto - функция суммирования элементов
	for _,i := range slice {
		sum += float64(i)
	}
	return
}

func findMaxaverage(num []int, k int) ( maxAvg float64 ) {
	// k - длина последовательно идущих элементов
	// нужно найти отрезок с максимальным средним числом
	n := len(num) - 1 // максимальный показатель индекса элемента
	maxAvg = -1000000
	
	
	for i:=0; i < n - k; i++ {
		if (Addition(num[i:i+k]) / float64(k) ) > float64(maxAvg) {

		}
	} 
	return
}

func main() {
	fmt.Println("Average max num for slice of array is: ", findMaxaverage())
}
