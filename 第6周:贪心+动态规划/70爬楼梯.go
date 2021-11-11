func climbStairs(n int) int {
    //ans := 0 
    dp := make([]int, n+1)
    dp[0]=0
   

    for i:=1; i<=n ;i++{

        if i==1 {
            dp[i]=1 
            continue
        }

        if i==2 {
            dp[i]=2
            continue 
        }

        dp[i] = dp[i-1]+dp[i-2]
    }

    return dp[n]


}


//作者：leo-ucf
//链接：https://leetcode-cn.com/problems/climbing-stairs/solution/70pa-lou-ti-dong-tai-gui-hua-by-leo-ucf-ayyr/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
