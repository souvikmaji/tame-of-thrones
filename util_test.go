package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountChar(t *testing.T) {
	tests := map[string]struct {
		input  string
		output map[rune]int
	}{
		"standard string": {
			input:  "hello",
			output: map[rune]int{'h': 1, 'e': 1, 'l': 2, 'o': 1},
		},
		"empty string": {
			input:  "",
			output: map[rune]int{},
		},
		"special character": {
			input:  "hello-world",
			output: map[rune]int{'h': 1, 'e': 1, 'l': 3, 'o': 2, '-': 1, 'w': 1, 'r': 1, 'd': 1},
		},
		"with space": {
			input:  "hello world",
			output: map[rune]int{'h': 1, 'e': 1, 'l': 3, 'o': 2, ' ': 1, 'w': 1, 'r': 1, 'd': 1},
		},
		"mixed case": {
			input:  "hello WORLD",
			output: map[rune]int{'h': 1, 'e': 1, 'l': 3, 'o': 2, ' ': 1, 'w': 1, 'r': 1, 'd': 1},
		},
	}

	for testName, test := range tests {
		t.Logf("Running test case %s", testName)
		assert.Equal(t, test.output, countChar(test.input))
	}

}
