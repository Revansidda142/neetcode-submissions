func topKFrequent(nums []int, k int) []int {
    freq := make(map[int]int)
    for _, v := range nums{
        freq[v]++
    }

    keys := []int{}
    for key, _ := range freq{
        keys = append(keys, key)
    }
    bubbleSortByFrequency(keys,freq)
    return keys[:k]
}

// Custom bubble sort for keys based on frequency
func bubbleSortByFrequency(keys []int, freq map[int]int) {
    n := len(keys)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if freq[keys[j]] < freq[keys[j+1]] {
                // Swap if next element is more frequent
                keys[j], keys[j+1] = keys[j+1], keys[j]
            }
        }
    }
}
