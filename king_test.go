package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewKing(t *testing.T) {
	tests := map[string]struct {
		input  string
		output *King
	}{
		"king shan": {
			input:  "shan",
			output: &King{Name: "shan", Kingdom: &Kingdom{}},
		},
		"king wth no name": {
			input:  "",
			output: &King{Name: "", Kingdom: &Kingdom{}},
		},
	}

	for testName, test := range tests {
		t.Logf("Running test case %s", testName)
		assert.Equal(t, test.output, newKing(test.input))
	}
}
