func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    m := make(map[rune]int)
    for _, val := range s{
        m[val]++
    }
    for _, val := range t{
        m[val]--
    }
    for _, count := range m {
        if count != 0 {
            return false
        }
    }
    return true
}
