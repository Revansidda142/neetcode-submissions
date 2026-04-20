func topKFrequent(nums []int, k int) []int {
    freq := make(map[int]int)
    for _, v := range nums{
        freq[v]++
    }

    keys := []int{}
    for key, _ := range freq{
        keys = append(keys, key)
    }
    sort.Slice(keys, func(i, j int) bool {
        return freq[keys[i]] > freq[keys[j]]
    })
    return keys[:k]
}
