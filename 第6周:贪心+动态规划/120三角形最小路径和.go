import (
    "fmt"
)

func minimumTotal(triangle [][]int) int {

    layer := len(triangle)
    ans := 100000
    end := 0 

   
    //记录上一个的位置
    pre := make([][][]int, layer)
    for i:=0;i<layer;i++{
        pre[i] = make([][]int, len(triangle[layer-1]))
    }

    //初始化位置为-1 -1 
    for i:=0;i<layer;i++{
        for j:=0;j<len(triangle[i]);j++{
         pre[i][j] = []int{-1,-1}

        }
    }
    //初始化DP
    dp := make([][]int, layer)
    for i:=0;i<layer;i++{
        dp[i] = make([]int, len(triangle[layer-1]))
    }
    dp[0][0]=triangle[0][0]

    //遍历DP

    for i:= 1;i<layer;i++{

        for j:=0;j<len(triangle[i]);j++{


          if j==len(triangle[i])-1{
              dp[i][j]=dp[i-1][j-1]+triangle[i][j]
              pre[i][j] = []int{i-1, j-1}
              continue
          }

           if j==0 {
              dp[i][j]=dp[i-1][0]+triangle[i][j]
               pre[i][j] = []int{i-1, 0}
              continue 
          }
    
          //转移方程
          if dp[i-1][j]>dp[i-1][j-1] {
              dp[i][j]=dp[i-1][j-1]+triangle[i][j]
              pre[i][j] = []int{i-1, j-1}
          }else {

              dp[i][j]=dp[i-1][j]+triangle[i][j]
              pre[i][j] = []int{i-1, j}

          }

        }


    }


    for i:=0;i<len(triangle[layer-1]);i++{

        if dp[layer-1][i]<ans{

            ans = dp[layer-1][i]
            end = i 
        }

        ans= min(ans, dp[layer-1][i])
        
    }
    
    //递归打印路径
    print(layer-1, end, pre, triangle)


    return ans 

}


func print(layer, end int, pre [][][]int, triangle [][]int) {

    if  pre[layer][end][0]!=-1 && pre[layer][end][1]!=-1 {
        print(pre[layer][end][0], pre[layer][end][1], pre, triangle)
    }

    fmt.Print("+")

    fmt.Print(triangle[layer][end])


}



func min (a, b int) int {
    if (a<b){return a}else{return b}
}

func max (a, b int) int {
    if (a>b){return a}else{return b}
}


//作者：leo-ucf
//链接：https://leetcode-cn.com/problems/triangle/solution/dpji-lu-zhuan-yi-lu-jing-by-leo-ucf-mdzi/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
