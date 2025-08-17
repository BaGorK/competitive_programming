class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        lp, rp = 0, 1
        maxP = 0
        while rp < len(prices):
            if prices[lp] < prices[rp]:
                profit = prices[rp] - prices[lp]
                maxP = max(maxP, profit)
            else:
                lp = rp
            rp+=1
        return maxP