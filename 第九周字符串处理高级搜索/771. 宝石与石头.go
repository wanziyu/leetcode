func numJewelsInStones(jewels string, stones string) int {
    jewelsCount := 0
    jewelsSet := map[byte]bool{}
    for i := range jewels {
        jewelsSet[jewels[i]] = true
    }
    for i := range stones {
        if jewelsSet[stones[i]] {
            jewelsCount++
        }
    }
    return jewelsCount
}
