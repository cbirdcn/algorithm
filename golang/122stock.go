package main

/**
122
贪心算法：买卖股票的最佳时机II
https://mp.weixin.qq.com/s/VsTFA6U96l18Wntjcg3fcg

给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:
输入: [7,1,5,3,6,4]
输出: 7
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4。随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。
*/

import "fmt"

func main() {
	prices1 := []int{7, 1, 5, 3, 6, 4}
	prices2 := []int{1, 2, 3, 4, 5}
	prices3 := []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices1))
	fmt.Println(maxProfit(prices2))
	fmt.Println(maxProfit(prices3))
}

// 贪心法（常规推断：收集每日正利润）
func maxProfit(prices []int) (profit int){
	for i:=0; i< len(prices) - 1; i++ {
		if growth := prices[i+1] - prices[i]; growth > 0 {
			profit += growth
		}
	}
	return profit
}

// 动态规划