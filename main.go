package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
	"unicode"

	"github.com/integrii/flaggy"
)

const (
	CASE_LOWER = iota
	CASE_SNAKE
	CASE_TITLE
	CASE_UPPER
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

func capitalize(word string) string {
	if word == "" {
		return word
	}

	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func generateCodename(words []string, count int, casetype int) string {
	if count < 1 {
		log.Fatal("word count must be at least 1")
	}

	if len(words) < count {
		log.Fatalf("need at least %d unique words to generate a codename", count)
	}

	indices := rand.Perm(len(words))[:count]
	var builder strings.Builder
	for i, idx := range indices {
		word := words[idx]
		switch casetype {
		case CASE_LOWER:
			word = strings.ToLower(word)
		case CASE_SNAKE:
			if i == 0 {
				word = strings.ToLower(word)
			} else {
				word = capitalize(word)
			}
		case CASE_TITLE:
			word = capitalize(word)
		case CASE_UPPER:
			word = strings.ToUpper(word)
		}
		builder.WriteString(word)
	}

	return builder.String()
}

func main() {
	var lowerCase bool
	var snakeCase bool
	var titleCase bool
	var upperCase bool

	wordCount := 2

	flaggy.Int(&wordCount, "c", "count", "Word count")
	flaggy.Bool(&lowerCase, "l", "lowercase", "Lowercase output")
	flaggy.Bool(&snakeCase, "s", "snakecase", "Snakecase output")
	flaggy.Bool(&titleCase, "t", "titlecase", "Titlecase output")
	flaggy.Bool(&upperCase, "u", "uppercase", "Uppercase output")
	flaggy.Parse()

	var casetype int
	if lowerCase {
		casetype = CASE_LOWER
	} else if snakeCase {
		casetype = CASE_SNAKE
	} else if titleCase {
		casetype = CASE_TITLE
	} else if upperCase {
		casetype = CASE_UPPER
	}

	fmt.Println(generateCodename(wordList, wordCount, casetype))
}
