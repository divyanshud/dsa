package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	freq := make([]int, 26)
	for _, char := range s {
		ascii := int(char)
		if ascii <= 90 {
			freq[ascii-65] += 1
		} else if ascii <= 122 {
			freq[ascii-97] += 1
		}
	}
	for _, char := range t {
		ascii := int(char)
		if ascii <= 90 {
			freq[ascii-65] -= 1
		} else if ascii <= 122 {
			freq[ascii-97] -= 1
		}
	}

	for _, count := range freq {
		if count > 0 {
			return false
		}
	}
	return true
}

func main() {
	s1 := "eat"
	s2 := "tea"
	s3 := "Hi"
	s4 := "Bye"
	s5 := "Ab"
	s6 := "mn"

	fmt.Println(fmt.Sprintf("Is \"%s\" and \"%s\" valid anagaram: %t", s1, s2, isAnagram(s1, s2)))
	fmt.Println(fmt.Sprintf("Is \"%s\" and \"%s\" valid anagaram: %t", s3, s4, isAnagram(s3, s4)))
	fmt.Println(fmt.Sprintf("Is \"%s\" and \"%s\" valid anagaram: %t", s5, s6, isAnagram(s5, s6)))
}
