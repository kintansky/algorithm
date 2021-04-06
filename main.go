package main

import "fmt"

type BinarySearch struct {
}

// 非递归的二分查找
func (b *BinarySearch) search(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// 递归实现
func (b *BinarySearch) search2(arr []int, left int, right int, target int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if arr[mid] == target {
		return mid
	}
	if arr[mid] > target {
		return b.search2(arr, left, mid-1, target)
	} else {
		return b.search2(arr, mid+1, right, target)
	}
}

func main() {
	var arr []int
	for i := 0; i < 10; i++ {
		arr = append(arr, i)
	}
	fmt.Println(arr)
	b := BinarySearch{}
	idx := b.search(arr, 8)
	fmt.Println(idx)
	idx = b.search2(arr, 0, len(arr)-1, 2)
	fmt.Println(idx)

}
