func longestPalindromeSubseq(s string) int {

     n := len(s)
     dp := make([][]int, n+2)

     if n==1{
         return 1
     }

     for i:=0;i<n+2;i++{
         dp[i] = make([]int, n+2)
         dp[i][i] = 1
     }

     newstring := " " + s 
     dp[0][0]=1
     
    for len :=1; len<=n; len++{

      for l := 1; l<=n-len;l++{

         r := l+len

         if newstring[l] == newstring[r]{

            dp[l][r]  =  dp[l+1][r-1]+2
         }else {
             dp[l][r] = max(dp[l+1][r], dp[l][r-1])
         }
        

      }

    }

    return dp[1][n]
     
}

func max(a, b int) int{
    if (a>=b){
        return a
    }else{
        return b 
    }
}
