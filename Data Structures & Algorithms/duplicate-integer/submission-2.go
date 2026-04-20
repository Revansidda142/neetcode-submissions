func hasDuplicate(nums []int) bool {
    m := make(map[int]string)
    for _, val := range nums {
        if m[val] == "true" {
            return true
        }
        m[val] = "true"
    }
    return false
}
