package main

import (
	"fmt"
	"strings"
)

func isPalindrome(word string) bool {
	// Convert the word to lowercase to ignore case sensitivity
	word = strings.ToLower(word)

	// Initialize two pointers, one from the start and one from the end of the word
	i := 0
	j := len(word) - 1

	// Iterate until the pointers meet
	for i < j {
		// If characters at the pointers are not equal, return false
		if word[i] != word[j] {
			return false
		}
		// Move the pointers towards each other
		i++
		j--
	}
	// If the loop completes without returning false, the word is a palindrome
	return true
}

func main() {
	// Define the words to check
	word1 := "madam"
	word2 := "text"

	// Check if the words are palindromes and print the result
	fmt.Printf("%s is a palindrome: %t\n", word1, isPalindrome(word1))
	fmt.Printf("%s is a palindrome: %t\n", word2, isPalindrome(word2))
}
