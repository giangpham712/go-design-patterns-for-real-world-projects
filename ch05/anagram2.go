package main

import (
	"errors"
	"fmt"
	"os"
)

func write(fname string, anagrams map[string][]string) {
	file, err := os.OpenFile(
		fname,
		os.O_WRONLY+os.O_CREATE+os.O_EXCL,
		0644,
	)
	if err != nil {
		msg := fmt.Sprintf(
			"Unable to create output file: %v", err,
		)
		panic(msg)
	}

	fmt.Println(file.Name())
}

func makeAnagrams(words []string, fname string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Failed to make anagram:", r)
		}
	}()

	anagrams := mapWords(words)
	write(fname, anagrams)
}

func mapWords(words []string) map[string][]string {
	return make(map[string][]string)
}

func loadFile(fname string) ([]string, error) {
	return nil, errors.New(
		"Dictionary file name cannot be empty.")
}

func main() {
	words, _ := loadFile("")

	makeAnagrams(words, "")
}
