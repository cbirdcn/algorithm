package main

import "fmt"

/**
https://mp.weixin.qq.com/s/X02S61WCYiCEhaik6VUpFA
题目：344. 反转字符串
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。

示例 1：
输入：["h","e","l","l","o"] 输出：["o","l","l","e","h"] 示例 2：
输入：["H","a","n","n","a","h"] 输出：["h","a","n","n","a","H"]

注意：题目要求原地修改字符串
*/

func main() {
	str1 := "hello"
	charStr1 := []byte(str1)
	reverseString(charStr1)
	fmt.Println(string(charStr1))

	str2 := "Hannah"
	charStr2 := []byte(str2)
	reverseString(charStr2)
	fmt.Println(string(charStr2))

}

// 原地修改字节数组，不需要返回
func reverseString(s []byte) {
	// 注意：golang不能写left++, right--，因为golang不支持逗号表示并列，所以要用平行赋值left, right = left + 1, right - 1
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		// 注意：golang不需要用tmp交换两值，直接平行赋值搞定
		s[left], s[right] = s[right], s[left]
	}
}
