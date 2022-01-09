package main

import (
	"fmt"
	"math"
)

/**
滑动窗口（双指针法）获取长度最小的子数组
https://mp.weixin.qq.com/s/UrZynlqi4QpyLlLhBPglyg

leetcode209

给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。

示例：
输入：s = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。

*/

func main(){
	s := 7
	nums := []int{2,3,1,2,4,3}
	//nums := []int{2, 3}
	minLength, start, end := minSubArrayLen(s, nums)
	if minLength != 0 {
		fmt.Println(fmt.Sprintf("min length is: %d, start number is nums[%d] = %d, end number is nums[%d] = %d", minLength, start, nums[start], end, nums[end]))
	} else {
		fmt.Println("got 0 sub array.")
	}
}

func minSubArrayLen(target int, nums []int) (int, int, int){
	res := math.MaxInt32 // 返回值，最终滑动窗口的长度（默认值max），如果比max小（表示存在子数组）就更新
	sum := 0 // 滑动窗口求和
	i := 0 // 窗口起始
	subLength := 0 // 窗口长度起始
	j := 0 // 窗口末尾

	for ; j< len(nums); j++ {
		sum += nums[j] // 固定左端，滑动右端，找到求和>=目标值的第一个子数组

		for ; sum >= target; { // while循环：固定右端，滑动左端，在查到的子数组中逐个排除掉左端，获得最短子数组长度
			subLength = j - i + 1 // 求得子数组的长度，备用
			if subLength < res {
				res = subLength // 当本次子数组长度比res短时，包括第一次比maxInt短，和后面循环时固定右滑动左找到更短子数组，这两种情况
			}
			sum -= nums[i] // 每次while循环中，依次排除左端第1个元素，判断剩余元素是否还能满足sum >= target，满足则获取到更短子数组。可以把i++写到这一行，sum -= nums[i++]，因为是先用i再右移。
			i++
		}
	}

	if res == math.MaxInt32 {
		return 0, -1, -1
	} else {
		return res, i - 1, j - 1 // 返回：满足条件的最小子数组长度，当时滑动窗口的左端（自增过），当时滑动窗口的右端（自增国）
	}

}
