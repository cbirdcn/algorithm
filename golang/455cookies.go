/*
455. 分发饼干
假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。

对每个孩子 i，都有一个胃口值 g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j，都有一个尺寸 s[j] 。如果 s[j] >= g[i]，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。

 
示例 1:

输入: g = [1,2,3], s = [1,1]
输出: 1
解释: 
你有三个孩子和两块小饼干，3个孩子的胃口值分别是：1,2,3。
虽然你有两块小饼干，由于他们的尺寸都是1，你只能让胃口值是1的孩子满足。
所以你应该输出1。
示例 2:

输入: g = [1,2], s = [1,2,3]
输出: 2
解释: 
你有两个孩子和三块小饼干，2个孩子的胃口值分别是1,2。
你拥有的饼干数量和尺寸都足以让所有孩子满足。
所以你应该输出2.
 

提示：

1 <= g.length <= 3 * 104
0 <= s.length <= 3 * 104
1 <= g[i], s[j] <= 231 - 1

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/assign-cookies
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	// 用例1
	g := []int{1, 2, 3}	// 小孩胃口
	s := []int{1, 1}	// 饼干尺寸

	// 用例2
	// g := []int{1, 2}	// 小孩胃口
	// s := []int{1, 2, 3}	// 饼干尺寸

	sort.Ints(g)
	sort.Ints(s)
	child := 0


	// 大饼干优先喂饱大胃口
	// 先从大到小遍历胃口，再遍历饼干，才能满足用例1。如果先饼干后胃口，会发现饼s遍历完了都没吃上饼干，而g中胃口=1的孩子还没轮到遍历就结束了。
	// j := len(s)-1
	// for i:=len(g)-1; i>=0; i-- {
	// 	// 将循环写成if+自减，可以避免一部分重复的循环。现在是n~n^2，嵌套循环是n^2
	// 	if j >= 0 && g[i] <= s[j] {
	// 		child++
	// 		j--
	// 	}
	// }

	// 小饼干优先喂饱小胃口
	// 先从小到大遍历饼干，再遍历胃口。让小饼干先喂饱小胃口
	// TODO

	// 大饼干优先喂饱大胃口（嵌套循环）
	// 先从大到小遍历饼干，在每个饼干被遍历时，从大到小循环嵌套遍历胃口。如果胃口正好匹配饼干，就胃口自减并break到饼干遍历过程，否则继续遍历胃口
	j:=len(g)-1
	for i:=len(s)-1; i>=0 && j>=0; i-- {
		for ;j>=0 && i>=0; j-- {
			if s[i] >= g[j] {
				child++
				fmt.Println(i, j)
				j--
				break	// break的存在以及j在外部初始化，使得复杂度不是O(n^2)，因为n个i并不是每个i都遍历了n个j，而是最多O(n^2)最小O(n)
			}
		}
	}

	fmt.Println(child)
}