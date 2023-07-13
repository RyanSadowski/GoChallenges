package main

import (
    "fmt"
    "strings"
    "regexp"
)

func main(){
    var sentence = "we love code"

    //trim it
    var trimmed = trimString(sentence)
    bytes := []byte(trimmed)
    // reverse the whole string
    reverseString(bytes)

    // reverse each word back to normal
    var start int = 0
    for i:=0; i < len(bytes); i ++ {
        if bytes[i] == ' ' {
            reverseString(bytes[start:i])
            start = i +1
        }

        // Skip consecutive spaces for edge case
        for start < len(bytes) && bytes[start] == ' ' {
            start++
        }
    }

    // reverse the last word
    reverseString((bytes[start:]))
    fmt.Println(string(bytes))

}

func reverseString(bytes []byte){
    var left int = 0
    var right int = len(bytes) -1

    for left < right {
        bytes[left], bytes[right] = bytes[right], bytes[left]
        left ++
        right --
    }
}

func trimString(str string) string {
	// Trim extra spaces at the beginning and end of the string
	str = strings.TrimSpace(str)

	// Replace multiple spaces with a single space
	regex := regexp.MustCompile("\\s+")
	str = regex.ReplaceAllString(str, " ")

	return str
}
