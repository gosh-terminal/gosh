package pkg


func Find(source []string, value string) bool {
    for _, item := range source {
        if item == value {
            return true
        }
    }
    return false
}
