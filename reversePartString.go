package main

import (
	"fmt"
	"math"
)

/**
题目：541. 反转字符串II
给定一个字符串 s 和一个整数 k，你需要对从字符串开头算起的每隔 2k 个字符的前 k 个字符进行反转。

如果剩余字符少于 k 个，则将剩余字符全部反转。

如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。

示例:
输入: s = "abcdefg", k = 2
输出: "bacdfeg"

分析：
其实在遍历字符串的过程中，只要让 i += (2 * k)，i 每次移动 2 * k 就可以了，然后判断是否需要有反转的区间。
因为要找的也就是每2 * k 区间的起点，这样写程序会高效很多。
*/

func main() {
	s := "abcdefg"
	k := 2
	fmt.Println(reversePartString(s, k))
}

func reversePartString(s string, k int) string {
	// golang中字符串不可修改，字节数组可以修改
	t := []byte(s)

	// 可以用len(s)作为结束，反正还没到结束时可能就不满足条件退出循环了
	for i := 0; i < len(s); i += (2 * k) {
		 // 判断再加k个字符有没有比字符串还长，注意math函数一般需要float64
		end := math.Min(float64(i+k), float64(len(s)))
		// 注意因为索引都是整数，所以可以从float64直接转到int，否则需要自己写min()函数
		subStr := t[i:int(end)]

		// 
		for m, n := 0, len(subStr)-1; m < len(subStr)/2; m, n = m+1, n-1 {
			subStr[m], subStr[n] = subStr[n], subStr[m]
		}
	}
	return string(t)
}
