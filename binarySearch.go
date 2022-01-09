package main

import "fmt"

/**
二分搜索插入位置
https://mp.weixin.qq.com/s/fCf5QbPDtE6SSlZ1yh_q8Q

leetcode编号35：搜索插入位置
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

示例 1:
输入: [1,3,5,6], 5
输出: 2

示例 2:
输入: [1,3,5,6], 2
输出: 1

示例 3:
输入: [1,3,5,6], 7
输出: 4

示例 4:
输入: [1,3,5,6], 0
输出: 0
*/

func main(){
	arr1 := []int{4} // 插入位置在0位置前
	arr2 := []int{1, 3} // 查找到已经存在的值
	arr3 := []int{1, 2, 4} // 插入到两数中间
	arr4 := []int{1, 2} // 插入到最后位置之后
	target := 3

	fmt.Println(binarySearch(arr1, target))
	fmt.Println(binarySearch(arr2, target))
	fmt.Println(binarySearch(arr3, target))
	fmt.Println(binarySearch(arr4, target))

}

func binarySearch(arr []int, target int) int{
	n := len(arr)
	if len(arr) == 0 {
		return 0
	}

	l := 0
	r := n - 1

	for ;l <= r; { // 循环，只要target在一个左闭右闭的区间[l, r]中
		m := (l + r) / 2 // 如果l+r是奇数，m落在floor(mid)上

		// 通过m位置的元素与target判断，改变l或r的位置，过滤掉一半数据
		if arr[m] < target {
			l = m + 1
		} else if arr[m] > target {
			r = m - 1
		} else {
			return m
		}
	}

	return r + 1
}
