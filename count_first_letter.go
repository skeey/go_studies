package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	words := os.Args[1:]

	counts := countLetters(words)

	show(counts)
}

func countLetters(words []string) map[string]int {
	counts := make(map[string]int)

	for _, word := range words {
		letter := strings.ToUpper(string(word[0]))
		count, findedLetter := counts[letter]
		
		if findedLetter {
			counts[letter] = count + 1
		} else {
			counts[letter] = 1
		}
	}

	return counts
}

func show(counts map[string]int) {
	fmt.Println("Count of words started in each letter:")
	
	for letter, count := range counts {
		fmt.Printf("%s = %d\n", letter, count)
	}
}
