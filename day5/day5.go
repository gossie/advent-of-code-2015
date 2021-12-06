package day5

import (
	"bufio"
	"os"
	"strings"
)

func readData(filename string) []string {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic("failed opening file")
	}

	words := make([]string, 0, 1000)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words
}

func NiceWordsOldRules(file string) int {
	count := 0
	for _, word := range readData(file) {
		if strings.Contains(word, "ab") || strings.Contains(word, "cd") || strings.Contains(word, "pq") || strings.Contains(word, "xy") {
			continue
		}

		numberOfVocals := strings.Count(word, "a") + strings.Count(word, "e") + strings.Count(word, "i") + strings.Count(word, "o") + strings.Count(word, "u")
		if numberOfVocals < 3 {
			continue
		}

		for index := range word {
			if index > 0 && word[index] == word[index-1] {
				count++
				break
			}
		}
	}

	return count
}

func NiceWordsNewRules(file string) int {
	count := 0
	for _, word := range readData(file) {
		// a pair of any two letters that appears at least twice in the string
		found := false
		for index := range word {
			if index > 0 && index < len(word)-2 && strings.Contains(word[index+1:], word[index-1:index+1]) {
				found = true
				break
			}
		}
		if !found {
			continue
		}

		// one letter which repeats with exactly one letter between them
		found = false
		for index := range word {
			if index > 1 && word[index] == word[index-2] {
				found = true
				break
			}
		}
		if !found {
			continue
		}

		count++
	}

	return count
}
