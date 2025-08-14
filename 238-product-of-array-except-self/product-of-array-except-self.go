func productExceptSelf(nums []int) []int {
    n := len(nums);
    res := make([]int, n)
    for i, _ := range res {
        res[i] = 1
    }

    for i := 1; i < n; i++ {
        res[i] = res[i - 1] * nums[i - 1]
    }

    postfix := 1;
    for i := len(nums) - 1; i >= 0; i-- {
        res[i] *= postfix;
        postfix *= nums[i]
    }

    return res
}