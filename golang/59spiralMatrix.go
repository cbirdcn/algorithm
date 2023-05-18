package main

import (
	"fmt"
	"math"
)

/*
题目59.螺旋矩阵II
给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3 输出: [ [ 1, 2, 3 ], [ 8, 9, 4 ], [ 7, 6, 5 ] ]
输入: 4 输出: [[1 2 3 4] [12 13 14 5] [11 16 15 6] [10 9 8 7]]
显示：
1	2	3
8	9	4
7	6	5
*/

func main() {

	n := 3 // 输入

	// 注意golang的二维数组初始化，使用make是因为var声明后直接赋值会造成运行时错误
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	makeSpiralMatrix(n, &res) // 传递引用比传递值更节省空间，注意函数的参数类型是指针
	fmt.Println(res)
}

func makeSpiralMatrix(n int, res *[][]int) {
	// 分析：n的平方个元素，按照顺时针向心螺旋填满nxn个格子，每行每列都是n个，起始位置为res[0][0]。需要模拟n/2圈，当n为奇数时中间要单独处理。
	// 按照左包含右不包含（左闭右开）来模拟填充每一行(列)的数据
	// 最上行从左到右：第1圈的从左到右就是固定x=0，y从0到<n-1；第2圈的从左到右就是固定x=1，y从1到<n-2；综上第c圈就是固定x=c-1,y从0到n-c;下一圈就是c++,x++
	// 最右列从上到下：固定y=n-1，x从0到<n-1
	// 最下行从右向左：固定x=n-1，y从n-1到>0
	// 最左列从下到上：固定y=0，x从n-1到>0
	// 这样四个左闭右开区间完成一圈，下一圈circle++，起点从[circle-1, circle-1]开始。n=奇数时，最后要+1中间格子

	startx, starty := 0, 0 // 每个圈的起始位置
	midxy := n / 2
	num := 1
	circle := 1

	// circle可以充当循环条件，circle<=n/2时就转一圈
	for circle <= n/2 {
		i, j := startx, starty
		for j = starty; j < n-circle; j++ {
			(*res)[startx][j] = num // (*res)[i][j]需要带上括号，否则会识别成*(res[i][j])
			num++
		}
		for i = starty; i < n-circle; i++ {
			(*res)[i][n-circle] = num
			num++
		}
		for j = n - circle; j > starty; j-- {
			(*res)[n-circle][j] = num
			num++
		}
		for i = n - circle; i > startx; i-- {
			(*res)[i][starty] = num
			num++
		}
		// 每循环完一圈后，圈数、起点xy都要自增
		circle++
		startx++
		starty++
	}
	// n=奇数时，需要单独处理中间格子，因为循环圈数只有n/2
	if n%2 == 1 {
		(*res)[midxy][midxy] = int(math.Pow(float64(n), 2)) // 注意：golang的pow需要float64格式
	}
}
