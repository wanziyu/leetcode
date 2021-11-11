func findNumberOfLIS(nums []int) int {

    n := len(nums)
    count := make([]int, n)
    ans := 0
    maxLen := 0 
    leng := make([]int, n)

    for i:=0;i<n;i++{

        count[i]=1
        leng[i]=1
        for j:=0;j<i;j++ {

            if nums[i]>nums[j]{

              if leng[i] < leng[j]+1{

               leng[i]=leng[j]+1
               count[i]=count[j]
              }else if leng[i]==leng[j]+1 {
                  count[i]+=count[j]

              }
    
            }
    }

         if leng[i] > maxLen {
            maxLen = leng[i]
            ans = count[i] // 重置计数
        } else if leng[i] == maxLen {
            ans +=count[i]
        }

 }

   return ans 
}
