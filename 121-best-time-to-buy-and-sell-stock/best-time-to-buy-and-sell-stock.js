/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function(prices) {
  let lp = 0;
  let rp = 1;
  let max = 0;

  while (rp < prices.length) {
    if (prices[rp] > prices[lp]) {
        const profit = prices[rp] - prices[lp];
        max = Math.max(profit, max);
    } else {
        lp = rp;
    }

    rp++;
  }  

  return max;
};