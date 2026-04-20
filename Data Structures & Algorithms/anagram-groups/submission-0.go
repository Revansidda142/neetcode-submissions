func groupAnagrams(strs []string) [][]string {
    m := make(map[string][]string)

    for _, str := range strs {
        s := []rune(str)
        bubbleSortRunes(s)

        key := string(s)
        m[key] = append(m[key], str)
    }

    var result [][]string
    for _, group := range m {
        result = append(result, group)
    }

    return result
}

func bubbleSortRunes(arr []rune) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}