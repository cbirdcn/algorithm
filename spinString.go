package main

import "fmt"

/**
题目：剑指Offer58-II.左旋转字符串
字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。

示例 1：
输入: s = "abcdefg", k = 2
输出: "cdefgab"

示例 2：
输入: s = "lrloseumgh", k = 6
输出: "umghlrlose"
  限制：
1 <= k < s.length <= 10000

*/

func main () {
	s1 := "abcdefg"
	k1 := 2
	t1 := []byte(s1)
	reversePart(t1, 0, k1 - 1)
	reversePart(t1, k1, len(s1) - 1)
	reversePart(t1, 0, len(s1) - 1)
	fmt.Println(string(t1))

	s2 := "lrloseumgh"
	k2 := 6
	t2 := []byte(s2)
	reversePart(t2, 0, k2 - 1)
	reversePart(t2, k2, len(s2) - 1)
	reversePart(t2, 0, len(s2) - 1)
	fmt.Println(string(t2))

}

// 对字节数组的任何一段反转（只接受[]byte类型）
func reversePart(t []byte, start int, end int) []byte {
	// a+b=奇数时，(a+b)/2是唯一的中间数，是否替换都无所谓
	// a+b=偶数时，(a+b)/2是中间两数的左数，必须和右数替换
	// 所以，<=(a+b)/2的数字都替换就完了。也可以写成<(a+b)/2+1
	for i, j := start, end; i <= (start + end) / 2; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
	return t
}