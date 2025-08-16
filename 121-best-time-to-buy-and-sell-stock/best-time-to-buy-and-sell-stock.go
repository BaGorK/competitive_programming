func maxProfit(prices []int) int {
    lp := 0
    rp := 1
    max := 0

    for rp < len(prices) {
        if prices[lp] < prices[rp] {
            profit := prices[rp] - prices[lp]
            if profit > max {
                max = profit
            }
        } else {
            lp = rp
        }
        rp++
    }

    return max
}