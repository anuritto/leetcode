package besttimetobuyandsellstock

func maxProfit(prices []int) int {
	profit,min := 0,prices[0]

	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if (prices[i] - min) > profit {
			profit = prices[i] - min
		}
	}
	return profit
}




// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/