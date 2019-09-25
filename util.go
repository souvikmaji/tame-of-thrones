package main

import "strings"

func emblem(name string) string {
	emblems := map[string]string{
		Land:  Panda,
		Water: Octopus,
		Ice:   Mammoth,
		Air:   Owl,
		Fire:  Dragon,
	}

	return emblems[strings.ToLower(name)]
}

func countChar(msg string) map[rune]int {
	countMap := make(map[rune]int)
	for _, char := range msg {
		countMap[char]++
	}

	return countMap
}
