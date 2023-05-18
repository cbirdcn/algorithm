/*
1. 两数之和
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。
示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

示例 2：
输入：nums = [3,2,4], target = 6
输出：[1,2]

示例 3：
输入：nums = [3,3], target = 6
输出：[0,1]
 

提示：
2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案
 

进阶：你可以想出一个时间复杂度小于 O(n2) 的算法吗？

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

分析：遍历数组，当map中不存在另一个值时，就把(当前值=>索引下标)存入map。直到找到另一个值和对应的索引下标。返回两个下标
*/

package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	// nums := []int{2, 7, 9, 11}

	fmt.Println(nums)

	target := 9
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int)

	for k1, v := range nums {
		if k2, ok := hashTable[target-v]; ok {
			return []int{k1, k2} // 注意：当k1对应的值不存在时先入map，第2次遍历k1时查到存在k2（上次存储的k1），此时k2<k1，如果要保证顺序，可以返回[]int{k2,k1}
		}
		hashTable[v] = k1
	}

	return nil
}