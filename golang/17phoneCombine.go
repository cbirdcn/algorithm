/*
17. 电话号码的字母组合
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例 1：
输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

示例 2：
输入：digits = ""
输出：[]

示例 3：
输入：digits = "2"
输出：["a","b","c"]
 

提示：
0 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/letter-combinations-of-a-phone-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

package main

import (
	"fmt"
)

var res [][]string
var path []string
//var dic map[int][]string

func main() {
	in := []int{2, 3, 4}
	dic := make(map[int][]string, 10)
	dic[0] = []string{}
	dic[1] = []string{}
	dic[2] = []string{"a", "b", "c"}
	dic[3] = []string{"d", "e", "f"}
	dic[4] = []string{"g", "h", "i"}
	dic[5] = []string{"j", "k", "l"}
	dic[6] = []string{"m", "n", "o"}
	dic[7] = []string{"p", "q", "r", "s"}
	dic[8] = []string{"t", "u", "v"}
	dic[9] = []string{"w", "x", "y", "z"}

	path = make([]string, 0)
	res = make([][]string, 0)
	starti := 0
	startj := 0
	backTracking(in, dic, starti, startj, path)
	fmt.Println(res)
}

func backTracking(in []int, dic map[int][]string, starti int, startj int, path []string) {
	if len(path) == len(in) {
		tmp := make([]string, len(in))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
//	for i:=0; i< len(in); i++ {
//		i := starti
//		strlen := dic[in[i]]
		for j:=startj; j<len(dic[in[starti]]); j++ {
			path = append(path, dic[in[starti]][j])
			backTracking(in, dic, starti+1, startj, path)
			path = path[:len(path)-1]
		}

//	}
}
