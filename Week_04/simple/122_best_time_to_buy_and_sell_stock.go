package simple

// 1 <= prices.length <= 3 * 10 ^ 4
// 0 <= prices[i] <= 10 ^ 4

// 输入股票价格列表，输出最大利润
// 一、决策问题，用回溯算法解决
// 1. 画决策树
// 2. 实现回溯算法
// 二、用贪心算法解决
// 1.买入后，只要有利可图，即可以卖掉
// 2.当天可以卖出再买入
func maxProfit(prices []int) int {
	// 参数检测

	profit := 0
	tmpProfit := 0

	for i, p := range prices[1:] {
		tmpProfit = p - prices[i]
		if tmpProfit > 0 {
			profit += tmpProfit
		}
	}

	return profit
}
