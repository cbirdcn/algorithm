/*
27. 移除元素
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

 

说明:

为什么返回数值是整数，但输出的答案是数组呢?

请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

你可以想象内部操作如下:
// nums 是以“引用”方式传递的。也就是说，不对实参作任何拷贝
int len = removeElement(nums, val);

// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
for (int i = 0; i < len; i++) {
    print(nums[i]);
}

示例 1：

输入：nums = [3,2,2,3], val = 3
输出：2, nums = [2,2]
解释：函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。你不需要考虑数组中超出新长度后面的元素。例如，函数返回的新长度为 2 ，而 nums = [2,2,3,3] 或 nums = [2,2,0,0]，也会被视作正确答案。
示例 2：

输入：nums = [0,1,2,2,3,0,4,2], val = 2
输出：5, nums = [0,1,4,0,3]
解释：函数应该返回新的长度 5, 并且 nums 中的前五个元素为 0, 1, 3, 0, 4。注意这五个元素可为任意顺序。你不需要考虑数组中超出新长度后面的元素。
 

提示：

0 <= nums.length <= 100
0 <= nums[i] <= 50
0 <= val <= 100


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/remove-element
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

package main

import "fmt"

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


// func main() {
// 	//	nums := []int{3,2,2,3}
// 	//	val := 3
// 	//	nums :=[]int{0,1,2,2,3,0,4,2}
// 	//	val := 2
// 		nums := []int{3}
// 		val := 3
	
// 		fastIndex := 0
// 		slowIndex := 0
	
// 		for ;fastIndex<=len(nums)-1; {
// 			if nums[fastIndex] != val {
// 				nums[slowIndex] = nums[fastIndex]
// 				slowIndex = slowIndex + 1
// 			}
// 			fastIndex = fastIndex + 1
// 		}
// 		fmt.Println(len(nums[:slowIndex]))
// 	}