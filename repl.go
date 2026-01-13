package main

import (
	"strings"
)

func cleanInput(text string) []string {
	result := strings.Fields(strings.ToLower(strings.TrimSpace(text)))
	return result
}
