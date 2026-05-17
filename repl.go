package main

import (
	"strings"
)

func clearInput(s string) []string {
	result := strings.Fields(strings.ToLower(s))
	return result
}
