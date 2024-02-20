package main

import (
	"fmt"
)

// create func to join array1 and array2
func merge(nums1, nums2 []int, m,n int) {
	pointer := m + n - 1 // Указатель на элемент финального массива
	pointerNum1 := m - 1 // Указатель на элемент первого массива
	pointerNum2 := n - 1 // Указатель на элемент второго массива

	for ; pointer >= 0; {
		if pointerNum2 < 0 {
			break
		}

		if pointerNum1 >= 0 && nums1[pointerNum1] > nums2[pointerNum2] {
			nums1[pointer] = nums1[pointerNum1]
			pointerNum1--
		}else {
			nums1[pointer] = nums2[pointerNum2]
			pointerNum2--
		}
		pointer--
	}
}
func countZero(n []int) int {
	c := 0
	for _,i := range n {
		if i == 0 { c += 1}
	}
	return c
}

func main() {
	// Arrays
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	m,n := len(nums1) - countZero(nums1), len(nums2)
	
	merge(nums1, nums2, m, n)
	fmt.Println(nums1)
}
