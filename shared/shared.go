package shared

import "os"

// StringInSlice checks if a string is present in a slice of strings.
func StringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

// PathExists checks if the given path exists.
func PathExists(path string) bool {
    _, err := os.Stat(path)
    return err == nil || !os.IsNotExist(err)
}
