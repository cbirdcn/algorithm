/*
75. 颜色分类

给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

必须在不使用库内置的 sort 函数的情况下解决这个问题。

// 对75题稍微改一下，让一组数据对3取模分别按照0,1,2的顺序重新排列数组，要求原地

示例 1：

输入：nums = [2,0,2,1,1,0]
输出：[0,0,1,1,2,2]
示例 2：

输入：nums = [2,0,1]
输出：[0,1,2]
 

提示：

n == nums.length
1 <= n <= 300
nums[i] 为 0、1 或 2
 

进阶：

你能想出一个仅使用常数空间的一趟扫描算法吗？

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/sort-colors
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func main() {
	// nums := []int{2, 0, 2, 1, 1, 0}
	nums := []int{5,7,8,1,0,9,1,2,5,12,0,2,45}
	k := 3

	n := len(nums)
	r := n - 1
	m := 1
	for l:=0; l<n && l<r; l++ {
		if nums[l] % k == 1 {
			for ; m<r; m++ {
				if nums[m] % k == 0 {
					tmp := nums[l]
					nums[l] = nums[m]
					nums[m] = tmp
					m++
					fmt.Println(nums)
					break
				}
			}
		}
		if nums[l] % k == 2 {
			for ;r>=l; r-- {
				if nums[r] % k == 0 {
					tmp := nums[l]
					nums[l] = nums[r]
					nums[r] = tmp
					r--
					fmt.Println(nums)
					break
				}
			}
		}
	fmt.Println(nums)
	}

	fmt.Println(nums)

}