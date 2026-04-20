func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    m := make(map[rune]int)
    for _, v := range s{
        m[v]++
    }
    for _, v := range t{
        m[v]--
    }
    for _, count := range m{
        if count != 0{
            return false
        }
    }
    return true
}

// Make an array
// loop on s string 