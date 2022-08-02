package utils

import (
	"regexp"
	"strings"
)

func SplitInputWord(inputWord string) []string {
	reg := regexp.MustCompile(`\\n`)
	inputWord = reg.ReplaceAllString(inputWord, "\n")
	reg = regexp.MustCompile(`\r\n`)
	inputWord = reg.ReplaceAllString(inputWord, "\n")

	return strings.Split(inputWord, "\n")
}
