package lib

import (
    "io/ioutil"
)

// Reads the file in "contents/" matching the supplied filename plus ".md"
func ReadContent(filename string) ([]byte, error) {
    filepath := "contents/" + ToFilename(filename)
    content, err := ioutil.ReadFile(filepath)
    return content, err
}
