/*
77. 组合
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

你可以按 任何顺序 返回答案。

示例 1：

输入：n = 4, k = 2
输出：
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
示例 2：

输入：n = 1, k = 1
输出：[[1]]
 

提示：

1 <= n <= 20
1 <= k <= n

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/combinations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

golang slice深拷贝：https://www.liangtian.me/post/go-slice/
*/

package main

import "fmt"

var result [][]int // 存放符合条件结果的集合
var path []int // 用来存放符合条件单一结果。path数组的大小如果达到k，说明我们找到了一个子集大小为k的组合了，在图中path存的就是根节点到叶子节点的路径。然后用result把path存起来，并终止本层递归。

func main () {
	n := 4
	k := 2
	fmt.Println(combine(n, k))
}

func combine(n int, k int) [][]int{
	backTracking(n, k ,1)
	return result
}

func backTracking (n int, k int, startIndex int) {
	// 递归终止条件
	if len(path) == k {
		// 注意：需要理解深拷贝。slice的拷贝是针对切片提供的接口，是深拷贝，可以通过调用copy()函数将src切片的值拷贝到dst切片中。通过该函数进行的切片拷贝后，针对dst切片进行的操作不会对src产生任何的影响，其拷贝长度是按照src与dst切片中最小的len长度去计算的。
		// 如果这里直接用result append path，就会发现for中第1行，当path经过一次回溯变成[1]再经过一次append()变成[1,3]时，result因为存的是path的深拷贝所以也跟着从[1,2]变成了[1,3]
		tmp := make([]int, k)
		copy(tmp, path)
		result = append(result, tmp) // 将节点组合加入结果集
		return
	}

	for ; startIndex<=n; startIndex++  {
		path = append(path, startIndex) // 加入节点
		backTracking(n, k, startIndex+1) // 递归
		path = path[:len(path)-1] // 回溯，撤销刚加入的节点
	}
}