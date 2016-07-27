package lib

import (
    "regexp"
    "strings"
)

// Converts input string `name` into a string with only lower case characters
// and dashes
func ToFilename(name string) string {
    r, _ := regexp.Compile("[^A-Za-z0-9 ]+")
    filename := r.ReplaceAllString(name, "")
    filename = strings.Trim(filename, " ")
    filename = strings.Replace(filename, " ", "-", -1)
    filename = strings.ToLower(filename)
    return filename + ".md"
}
