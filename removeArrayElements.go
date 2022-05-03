package main

import (
	"fmt"
)

/*
编号：27. 移除元素
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并「原地」修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

示例 1:
给定 nums = [3,2,2,3], val = 3,
函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。
你不需要考虑数组中超出新长度后面的元素。

示例 2:
给定 nums = [0,1,2,2,3,0,4,2], val = 2,
函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。
*/

func main() {
	// 数组写法1：数组类型为"[数量]type"，元素的数量 n 也是该类型的一部分。
	// 所以如果定义了数量就只能让函数接收固定个数的数组才行。也就是函数形参为nums [4]int，不适合作为函数参数
	// nums1 := [4]int{3,2,2,3}
	// nums2 := [8]int{0,1,2,2,3,0,4,2}

	// 数组写法2：...意思是忽略长度，由编译器计算数组长度，但是本质上和写法1相同。不适合作为函数参数
	// nums1 := [...]int{3,2,2,3}
	// nums2 := [...]int{0,1,2,2,3,0,4,2}

	// 声明切片：切片是创建了一个数组，并引用了这个数组。所以不需要提供长度，编译器可以计算。可以作为函数参数。
	nums1 := []int{3, 2, 2, 3}
	nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	nums3 := []int{3, 2, 2}

	target := 2
	fmt.Println(removeArrayElements(nums1, target))
	fmt.Println(removeArrayElements(nums2, target))
	fmt.Println(removeArrayElements(nums3, target))
}

// 快慢指针法
func removeArrayElements(nums []int, target int) int {
	// 两个指针都从头开始，l慢r快
	l, r := 0, 0
	for ; r < len(nums); r++ {
		if nums[r] != target {
			l++
		}
	}
	// 经过循环中最后一次不是目标值的判断后，l+1，落到了下一个位置，所以直接返回l而不是l+1
	return l
}

// 头尾指针法
func removeArrayElements2(nums []int, target int) int {
	// // 两个指针从头(0)、尾的下一位(len)向中间走
	// l、r分别定位在数组首位和最后一位的下一位。
	// 这样for l < r，最后l+1=r，此时如果相等对应值返回l(不再加最后一个值)，如果不等l++添加最后一个值还是返回l
	// 反之，假设r在末尾位置(len-1)，for l<=r，最后l=r，此时匹配要返回l+1，不匹配要返回l。
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == target {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}
