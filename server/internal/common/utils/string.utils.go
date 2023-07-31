package utils

import (
	"regexp"
	"strings"
)

func RemoveSpecialChars(s string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	return reg.ReplaceAllString(s, "-")
}

func GenerateAliasString(str string) string {
	// Convert to lowercase
	lowercaseInput := strings.ToLower(str)
	// Replace spaces with dashes
	dashedInput := strings.ReplaceAll(lowercaseInput, " ", "-")
	// Remove special characters and keep only alphanumeric characters and dashes
	finalOutput := RemoveSpecialChars(dashedInput)

	return finalOutput
}
