package main

import(
	"fmt"
	"math"
)

/**
202 快乐数
https://mp.weixin.qq.com/s/G4Q2Zfpfe706gLK7HpZHpA

编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。如果 可以变为  1，那么这个数就是快乐数。

如果 n 是快乐数就返回 True ；不是，则返回 False 。

「示例：」

输入：19
输出：true
解释：
1^2 + 9^2 = 82
8^2 + 2^2 = 68
6^2 + 8^2 = 100
1^2 + 0^2 + 0^2 = 1

题目中说了会 「无限循环」，那么也就是说「求和的过程中，sum会重复出现，这对解题很重要！」

「当我们遇到了要快速判断一个元素是否出现集合里的时候，就要考虑哈希法了。」

「还有一个难点就是求和的过程，如果对取数值各个位上的单数操作不熟悉的话，做这道题也会比较艰难。」

*/

func getSum(n int) int{
	s := 0

	// 获取单数，就是循环对10取余获得尾数，再整除10获取去掉尾数的前数，直到整除后=0
	for n > 0 {
		s += int(math.Pow(float64(n % 10), 2)) // 传入int，但是math操作需要float64，power计算完还要转回int用于下一轮循环
		n /= 10
	}

	return s
}

func main() {
	n := 19
	set := make(map[int]struct{})

	for {
		n = getSum(n)
		fmt.Println(n)
		if n != 1 {
			// // 如果这个sum曾经出现过，说明已经陷入了无限循环了，立刻return false
			if _, ok := set[n]; ok {
				fmt.Println(false)
			} else {
				set[n] = struct{}{} // 不占空间
			}
		} else {
			break
		}
	}

	fmt.Println(true)
}