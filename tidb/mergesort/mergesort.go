package main

import "fmt"
// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func Sort(nums []int64) []int64 {
	//defer fmt.Println(nums)
	if len(nums) <= 1 {
		return nums
	}
	leng := int64(len(nums))
	mid := (leng - 1) / 2
	new_arr := make([]int64, leng)
	left := Sort(nums[:mid + 1])
	right := Sort(nums[mid + 1:])

	var i, j, k int64
	// fmt.Println(i < mid, j < (leng - mid))

	for i < int64(len(left)) && j < int64(len(right)) {
		if left[i] < right[j] {
			new_arr[k] = left[i]
			i++
		} else {
			new_arr[k] = right[j]
			j++
		}
		k++
	}
	// defer fmt.Println(nums, left, right, leng, i, j, k)

	if i == int64(len(left)) {
		for ; j < int64(len(right)); j++ {
			new_arr[k] = right[j]
			k++
		}
	} else {
		for ; i < int64(len(left)); i++ {
			new_arr[k] = left[i]
			k++
		}
	}
	return new_arr
}

func MergeSort(src []int64) {
	fmt.Println(Sort(src))
}

func main() {
	arr := []int64{2,3,5,1,3,4,6,4,2}
	MergeSort(arr)
}