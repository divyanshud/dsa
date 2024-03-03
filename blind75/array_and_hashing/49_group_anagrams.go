package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]uint8][]string)
	for _, str := range strs {
		var freq [26]uint8
		for _, char := range str {
			ascii := int(char)
			freq[ascii-97] += 1
		}

		groups[freq] = append(groups[freq], str)
	}
	results := make([][]string, 0)
	for _, v := range groups {
		results = append(results, v)
	}
	return results
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println("The similar anagrams are", groupAnagrams(strs))
}
