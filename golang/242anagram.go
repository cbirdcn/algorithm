/*
242. 有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

 

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false
 

提示:

1 <= s.length, t.length <= 5 * 104
s 和 t 仅包含小写字母
 

进阶: 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/valid-anagram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

分析：
定义一个数组叫做record用来上记录字符串s里字符出现的次数。
需要把字符映射到数组也就是哈希表的索引下表上，「因为字符a到字符z的ASCII是26个连续的数值，所以字符a映射为下表0，相应的字符z映射为下表25。」
再遍历字符串s的时候，「只需要将 s[i] - ‘a’ 所在的元素做+1 操作即可，并不需要记住字符a的ASCII，只要求出一个相对数值就可以了。」 这样就将字符串s中字符出现的次数，统计出来了。
那看一下如何检查字符串t中是否出现了这些字符，同样在遍历字符串t的时候，对t中出现的字符映射哈希表索引上的数值再做-1的操作。
那么最后检查一下，「record数组如果有的元素不为零0，说明字符串s和t一定是谁多了字符或者谁少了字符，return false。」
最后如果record数组所有元素都为零0，说明字符串s和t是字母异位词，return true。
golang可以用map代替数组，底层使用哈希实现。
*/
package main

import "fmt"

func main() {
	s1 := "anagram"
	t1 := "nagaram"
	s2 := "rat"
	t2 := "car"
	isAnagram(s1, t1)
	isAnagram(s2, t2)
}

func isAnagram(s string, t string) (bool){
	record := make(map[int32]int, 26)
	for _, v := range s {
		// fmt.Println(reflect.TypeOf(v))
		// fmt.Println(v)         // range 字符串后得到int32类型的数字
		// fmt.Println(string(v)) // 可以将int32强制转成string
		record[v]++
	}
	for _, v := range t {
		record[v]--
	}
	for k, v := range record {
		if v != 0 {
			fmt.Printf("not anagram, character: %s \n", string(k))
			return false
		}
	}
	fmt.Printf("is anagram \n")
	return true
}