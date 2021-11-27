func findRedundantConnection(edges [][]int) []int {
    
    n := len(edges)

    //初始化并查集
    fa := make([]int, n+1)

    for i:=0;i<=n;i++{

        fa[i] = i

    }

    for _, edge := range edges {

        x := edge[0]
        y := edge[1]

        //如果两个元素的根是一个说明 已经成环了 直接返回即可
        if (Find(fa,x)==Find(fa,y)){

            return []int{x, y}
        }else{
            unionSet(fa, x, y)
        }


    }

    return []int{}

}


//并查集函数Find
func Find(fa []int,  x int) int {
    if (fa[x]==x){return  x}
     fa[x] = Find(fa, fa[x])

     return fa[x]
}  
//并查集函数unionSet
func unionSet(fa []int,  x, y int ) {
    x, y = Find(fa, x), Find(fa, y) 
    if (x !=y){
        fa[x] = y
    }
}
