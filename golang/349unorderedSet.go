package main

/**
无序集合
https://mp.weixin.qq.com/s/N9iqAchXreSVW7zXUS4BVA

第349题. 两个数组的交集
题意：给定两个数组，编写一个函数来计算它们的交集。
「说明：」
输出结果中的每个元素一定是唯一的。
我们可以不考虑输出结果的顺序。
分析：
这道题目没有限制数值的大小，就无法使用数组来做哈希表了。
如果哈希值比较少、特别分散、跨度非常大，使用数组就造成空间的极大浪费！
题目特意说明：「输出结果中的每个元素一定是唯一的，也就是说输出的结果的去重的， 同时可以不考虑输出结果的顺序」
*/

import "fmt"

func main() {
	nums1 := []int{1,2,2,1}
	nums2 := []int{2,2}
	nums3 := []int{4,9,5}
	nums4 := []int{9,4,9,8,4}
	fmt.Println(getIntersection(nums1, nums2))
	fmt.Println(getIntersection(nums3, nums4))
}

func getIntersection(nums1 []int, nums2 []int) (intersection []int) {
	set1 := make(map[int]struct{})
	set2 := make(map[int]struct{})

	for _,v := range nums1 {
		set1[v] = struct{}{} // struct{}{}是struct{}空类型且空值的不占空间的value，非常适合将map改装成set。
	}
	for _,v := range nums2 {
		set2[v] = struct{}{}
	}

	// 让nums1是更短的一个，遍历耗时少
	if len(nums1) > len(nums2) {
		set1, set2 = set2, set1
	}

	// 遍历更短的set1
	for k := range set1 {
		// set2存在对应元素
		if _, ok := set2[k]; ok {
			intersection = append(intersection, k)
		}
	}
	return
}