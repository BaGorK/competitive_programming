/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
  const prevMark = new Map()
  for (let i = 0; i < nums.length; i++) {
    const diff = target - nums[i];
    if (prevMark.has(diff)) {
        return [prevMark.get(diff), i] 
    }
    prevMark.set(nums[i], i)
  }
  return []
};