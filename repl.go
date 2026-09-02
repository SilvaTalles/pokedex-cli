package main

import (
	"strings"
)

func cleanInput(text string) []string {
	var result []string
	text = strings.ToLower(text)
	text = strings.TrimSpace(text)
	i := strings.Index(text, " ")
	for i > -1 {
		result = append(result, text[:i])
		text = text[i+1:]
		i = strings.Index(text, " ")
	}
	result = append(result, text)
	return result
}
