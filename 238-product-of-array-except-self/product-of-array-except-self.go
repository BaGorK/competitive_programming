func productExceptSelf(nums []int) []int {
    n := len(nums);
    res := make([]int, n)
    for i, _ := range res {
        res[i] = 1
    }

    prefix := 1
    for i := 0; i < n; i++ {
        res[i] = prefix
        prefix *= nums[i]
    }
    postfix := 1;
    for i := len(nums) - 1; i >= 0; i-- {
        res[i] *= postfix;
        postfix *= nums[i]
    }

    return res
}