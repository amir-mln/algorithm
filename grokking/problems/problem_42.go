package problems

import "math"

/*
# Best Time to Buy and Sell Stock

Given an array where the element at the index [i] represents
the price of a stock on day [i], find the maximum profit that
you can gain by buying the stock once and then selling it.

You can buy the stock on one specific day and sell it on a
different day to make a profit. If no profit can be achieved,
we return zero.

Constraints:

 1. We can’t sell before buying a stock, that is, the array
    index at which stock is bought will always be less than the
    index at which the stock is sold.
 2. 1 ≤ len(Problem42) ≤ 1000
 3. 0 ≤ Problem42[i] ≤ 100_000

Description:

We need to find the two prices that have the maximum difference.
Therefore, we need a variable for tracking the minimum price and
another for tracking the maximum profit. As we iterate through
the prices, we first update the minimum price variable. Then, we
subtract the current price from the minimum price. Finally, if
the result of this subtraction is larger than the maximum profit
variable, we'll update it.
*/
type Problem42 []int

func (p Problem42) Solve() int {
	minPrice := math.MaxInt
	maxProf := math.MinInt

	for _, p := range p {
		minPrice = min(minPrice, p)
		maxProf = max(maxProf, p-minPrice)
	}

	return maxProf
}
