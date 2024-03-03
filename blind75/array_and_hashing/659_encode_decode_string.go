package main

import (
	"fmt"
	"strconv"
	"strings"
)

func encode(strs []string) string {
	results := strings.Builder{}
	for idx := range strs {
		results.WriteString(fmt.Sprintf("%d#%s", len(strs[idx]), strs[idx]))
	}
	return results.String()
}

func decode(str string) []string {
	results := make([]string, 0)
	size := len(str)
	if size >= 0 {
		i := 0
		for i < size {
			j := i
			for str[j] != '#' {
				j++
			}
			length, _ := strconv.Atoi(str[i:j])
			results = append(results, str[j+1:(j+1+length)])
			i = j + 1 + length
		}
	}
	return results
}

func main() {
	strs := []string{"Hello", "World"}
	encodedStr := encode(strs)
	fmt.Println(encodedStr)
	decodedStrs := decode(encodedStr)
	isValid := false
	if len(strs) == len(decodedStrs) {
		for i := 0; i < len(strs); i++ {
			if strs[i] == decodedStrs[i] {
				isValid = true
			}
		}
	}
	fmt.Println(fmt.Sprintf("Is Valid: %t", isValid))
}
