package main

import (
	"fmt"
	// "golang.org/x/text/unicode/rangetable"
)

/**
第1题. 两数之和
https://mp.weixin.qq.com/s/uVAtjOHSeqymV8FeQbliJQ

给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

「示例:」

给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]


暴力的解法是两层for循环查找，时间复杂度是O(n^2)。

两数之和这道题目，不仅要判断y是否存在而且还要记录y的下表位置，因为要返回x 和 y的下表。所以set 也不能用。

使用哈希法最为合适

map是一种key value的存储结构，可以用key保存数值，用value在保存数值所在的下标。

分析：遍历数组，当map中不存在另一个值时，就把(当前值=>索引下标)存入map。直到找到另一个值和对应的索引下标。返回两个下标
*/

func main() {
	nums := []int{2, 7, 11, 15}

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
