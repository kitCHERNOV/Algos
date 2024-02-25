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
		// Если элемент подмассива окажетя пустым, то каретку нужно будет сдвинуть на значение +n где n это смещение по head
		if &num[i] == nil{
			continue
		}else if &num[i+1] == nil {
			i++
			continue
		}else if &num[i+2] == nil {
			i+=2
			continue
		}else if &num[i+3] == nil {
			i+=3
			continue
		}
		// s := Addition(num[i:i+k])
		// fmt.Println(s, s/float64(k))
		if s := Addition(num[i:i+k]); ( s / float64(k) ) > maxAvg {
			maxAvg = s / float64(k)
		}

	} 
	return
}

func main() {
	const length int = 10
	var nums = [length]int{1,2,3,4,5}
	var lforslice int = 4
	fmt.Println("Average max num for slice of array is: ", findMaxaverage(nums[:], lforslice))
}
