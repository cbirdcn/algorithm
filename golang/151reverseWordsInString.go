package main

import "fmt"

/**
https://mp.weixin.qq.com/s/X3qpi2v5RSp08mO-W5Vicw
题目：151.翻转字符串里的单词
给定一个字符串，逐个翻转字符串中的每个单词。

示例 1：
输入: "the sky is blue"
输出: "blue is sky the"

示例 2：
输入: "  hello world!  "
输出: "world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。

示例 3：
输入: "a good   example"
输出: "example good a"
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

分析：
假设不允许使用辅助空间
我们将整个字符串都反转过来，那么单词的顺序指定是倒序了，只不过单词本身也倒叙了，那么再把单词反转一下，单词就正过来了。
所以解题思路如下：
	移除多余空格
	将整个字符串反转
	将每个单词反转
*/

func main() {
	s1 := "the sky is blue"
	s2 := "  hello world!  "
	s3 := "a good   example"

	fmt.Println(reverseWordsInString(reverseAll(trimSpacesInString(s1))))
	fmt.Println(reverseWordsInString(reverseAll(trimSpacesInString(s2))))
	fmt.Println(reverseWordsInString(reverseAll(trimSpacesInString(s3))))
}

func trimSpacesInString(s string) string {
	t := []byte(s)
	newEnd := 0

	// 双指针获取去掉空格后的字符串(i原指针，j新指针)
	for i := 0; i < len(t) && newEnd <= i; i, newEnd = i+1, newEnd+1 {
		if t[i] == ' ' {
			if newEnd == 0 && t[i] == ' ' {
				// 开头多空格
				newEnd--
			} else if i < len(t)-1 {
				// 中间多空格
				if t[i+1] == ' ' {
					newEnd--
				} else {
					t[newEnd] = t[i]
				}
			} else {
				// 末尾多空格
				newEnd--
				continue
			}
		} else {
			t[newEnd] = t[i]
		}
	}

	// newEnd作为函数内有效的变量，for循环后被自动+1，也就是newEnd是实际能取到末尾的下一位。但是因为slice切片含前不含后，所以直接用newEnd作为第二个参数
	return string(t[0:newEnd])
}

func reverseAll(s string) string {
	t := []byte(s)
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
	return string(t)
}

func reverseWordsInString(s string) string {
	t := []byte(s)
	for slow, fast := 0, 0; fast < len(s); fast++ {
		// 注意末尾没有空格，要单独处理。末尾不能用+1位置的索引，会导致slice越界
		if (fast < len(s)-1 && t[fast+1] == ' ') || (fast == len(s)-1) {
			// 从单词头到单词中间循环替换（首位指针向中间收敛到一半）
			for i, j := slow, fast; i < (slow+fast)/2+1; i, j = i+1, j-1 {
				t[i], t[j] = t[j], t[i]
			}
			slow = fast + 2 // slow要调整到新位置:fast+1空格+1，到下一个单词的起始位置
		}
	}

	return string(t)
}
