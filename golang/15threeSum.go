package main

import (
	"fmt"
	"sort"
)

/**
https://mp.weixin.qq.com/s/r5cgZFu0tv4grBAexdcd8A
https://leetcode-cn.com/problems/3sum/solution/san-shu-zhi-he-by-leetcode-solution/
第15题. 三数之和
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：[ [-1, 0, 1], [-1, -1, 2] ]
*/

func main() {
	nums1 := []int{-1, 0, 1, 2, -1, -4}
	nums2 := []int{-1, -1, 0, 0, 1, 1}
	nums3 := []int{-2, -1, -1, 0, 0, 0, 1, 1, 2}
	nums4 := []int{0, 0}
	fmt.Println(threeSum(nums1))
	fmt.Println(threeSum(nums2))
	fmt.Println(threeSum(nums3))
	fmt.Println(threeSum(nums4))
	/*
		[[-1 -1 2] [-1 0 1]]
		[[-1 0 1]]
		[[-2 0 2] [-2 1 1] [-1 -1 2] [-1 0 1] [0 0 0]]
		[]
	*/
}

func threeSum(nums []int) [][]int {
	// 对数字升序排列，得到[-4,-1,-1,0,1,2]
	sort.Ints(nums)
	// 结果集是二维slice
	ans := make([][]int, 0)
	n := len(nums)

	// 分析：从左向右枚举一个数字a（取索引值作为first指针）
	// 后面的数字索引作为剩下两个数字的选池。
	// 其中second作为左指针（初始first+1），third作为右指针(初始n-1)，分别向中间走找到second+third=-first的情况或左右指针指向同一个数字截止
	for first := 0; first < n; first++ {
		// 如果两次first枚举了同样的数字，则跳过。比如-1,-1...，由于题目要求不重复的三元组，则第二个-1直接跳过
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// 如果枚举的a>0，则b>a>0,c>a>0，所以a+b+c>0，直接退出返回结果集
		if nums[first] > 0 {
			return ans
		}
		// 每次枚举一个first指针后，都将third指针的初始位置指向最后
		third := n - 1
		// second + third 的目标值
		target := -1 * nums[first]

		// 继续枚举数字b（second）
		for second := first + 1; second < third; second++ {
			// second也不重复
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 如果b+c>-a，则c应该一直左移，直到不满足条件second<third或找到target
			// 由于前面for循环中规定second<third，所以当second+1=third时，执行完third--。然后for判断second==third就退出循环，不能让third--了，否则将会出现third+1=second的情况
			for (second != third) && (nums[second]+nums[third] > target) {
				third--
			}
			// 当second==third时，此first已经找不到对应的second-third对了，进入下一个first循环。由于上面for已经阻碍了出现third<second的情况，所以不需要second>=third才break
			if second == third {
				break
			}
			// 如果找到target
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}

	}

	return ans
}
