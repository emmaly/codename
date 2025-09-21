package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

//go:embed words.txt
var wordData string

var wordList []string

func init() {
	wordList = parseWords(wordData)
	if len(wordList) < 100 {
		log.Fatalf("expected at least 100 words, got %d", len(wordList))
	}

	rand.Seed(time.Now().UnixNano())
}

func parseWords(data string) []string {
	scanner := bufio.NewScanner(strings.NewReader(data))
	var words []string
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word == "" {
			continue
		}
		words = append(words, strings.ToLower(word))
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to read embedded word list: %v", err)
	}

	if len(words) == 0 {
		log.Fatal("embedded word list is empty")
	}

	return words
}

func generateCodename(words []string) string {
	if len(words) < 2 {
		log.Fatal("need at least two words to generate a codename")
	}

	firstIndex := rand.Intn(len(words))
	secondIndex := rand.Intn(len(words) - 1)
	if secondIndex >= firstIndex {
		secondIndex++
	}

	return words[firstIndex] + words[secondIndex]
}

func main() {
	fmt.Println(generateCodename(wordList))
}
