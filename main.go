package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {

	cleaned := []string{}

	split_list := strings.Split(text, " ")

	for _, v := range split_list {
		if strings.TrimSpace(v) != "" {
			cleaned = append(cleaned, strings.ToLower(strings.TrimSpace(v)))
		}

	}

	return cleaned
}
