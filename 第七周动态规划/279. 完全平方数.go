
func numSquares(n int) int {


   var len int = 0
   for i:=1;i*i<=n;i++{
     len = i
   }
   
   dp := make([]int, n+1)

   for i:=0;i<=n;i++{
       dp[i]=1000000
   }

   dp[0]=0

   for i:=1;i<=len+1;i++{

       for j:=i*i;j<=n;j++{
           dp[j]=min(dp[j], dp[j-i*i]+1)
       }
   }

   return dp[n]
}

func min(a, b int) int{
    if (a>=b){
        return b
    }else{
        return a 
    }
}
