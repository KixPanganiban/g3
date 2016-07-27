package lib

import (
    "fmt"
    "io/ioutil"
)

// Reads the file in "contents/" matching the supplied filename plus ".md"
func ReadContent(filename string) ([]byte, error) {
    filepath := "contents/" + ToFilename(filename)
    fmt.Println(filepath)
    content, err := ioutil.ReadFile(filepath)
    return content, err
}
