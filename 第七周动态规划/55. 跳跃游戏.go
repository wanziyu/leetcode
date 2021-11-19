func canJump(nums []int) bool {
    if len(nums)<=1{
        return true
    }
    dp:=make([]bool,len(nums))
    // 第一个肯定能到达(起始条件)
    dp[0]=true
    for i:=1;i<len(nums);i++{
        for j:=i-1;j>=0;j--{
            if dp[j] && nums[j]>=i-j{
                dp[i]=true
                break
            }
        }
    }
    return dp[len(nums)-1]
}
