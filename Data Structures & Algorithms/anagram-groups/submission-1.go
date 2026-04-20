func groupAnagrams(strs []string) [][]string {
    m := make(map[string][]string)

    for _, str := range strs {
        s := []rune(str)
        sort.Slice(s, func( i, j int) bool {
            return s[i] < s[j]
        })

        key := string(s)
        m[key] = append(m[key], str)
    }

    var result [][]string
    for _, group := range m {
        result = append(result, group)
    }

    return result
}