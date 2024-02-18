package main

import (
	"fmt"
	"math/rand"
)

func quickSort(ar []int) {
	if len(ar) < 2 {
		return
	}

	left, right := 0, len(ar) - 1
	pivotIndex := rand.Int() % len(ar)

	ar[pivotIndex], ar[right] = ar[right], ar[pivotIndex]

	for i := 0; i < len(ar); i++ {
		if ar[i] < ar[right] {
			ar[i], ar[left] = ar[left], ar[i]
			left++
		}
	}

	ar[left], ar[right] = ar[right], ar[left]

	quickSort(ar[:left])
	quickSort(ar[left + 1:])

	return
}

// create func to join array1 and array2
func merge(num1, nums2 []int) {
	for k,i := range nums2 {
		num1[k+3] = i
	}
	quickSort(num1)
}

func main() {
	// Arrays
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	
	merge(nums1, nums2)
	fmt.Println(nums1)
}
